#include <windows.h>
#include <stdlib.h>

/**
 * Extracted from https://www.youtube.com/watch?v=RiweaH6Qmro
*/

#define WINDOW_WIDTH 900
#define WINDOW_HEIGHT 600
#define UWM_UPDATEPOSITIONS (WM_USER + 1)
#define IDT_TIMER1 1
#define RADIUS 20



LRESULT WinProc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam)
{
        static int x = WINDOW_WIDTH/2 , y = WINDOW_HEIGHT/2;
        static int speed_x = 3;
        static int speed_y = 3;
        static RECT rect;

        switch (uMsg) {
        case WM_TIMER:
        {
            if (wParam == IDT_TIMER1){
                PostMessage(hwnd, UWM_UPDATEPOSITIONS, 0, 0);
            }
        }
        case UWM_UPDATEPOSITIONS:
        {
            /**
            * Here we get the rect because the window's height and width are not equivalent to the visible portion of it.
            * Rect contains information about the height, width and its margins.
            */
            GetClientRect(hwnd, &rect);
            x += speed_x;
            y += speed_y;
            /**
            * Important to not use the WINDOW_WIDTH macro instead of the rect.right as this will cause issues with the bottom and right margins.
            */
            if (x + RADIUS >= rect.right) {
                speed_x *= -1;
                x = rect.right - RADIUS;
            }
            else if(x - RADIUS <= 0){
                speed_x *= -1;
                x = RADIUS;
            }
            if (y + RADIUS >= rect.bottom) {
                speed_y *= -1;
                y = rect.bottom - RADIUS;
            }
            else if (y - RADIUS <= 0) {
                speed_y *= -1;
                y = RADIUS;
            }
            /*
            * With this method we force the OS to reload the page and, so, to re-paint the new items.
            * To sum up, this forces a WM_PAINT message
            */
            InvalidateRect(hwnd, NULL, TRUE);
        }
        case WM_CREATE:
        {
            SetTimer(hwnd, IDT_TIMER1, 1000 / 60, NULL);  // 60 cops per segon (aprox. 16.67 ms per tick)
            break;
        }
        case WM_PAINT:
        {
            PAINTSTRUCT ps;
            HDC hdc = BeginPaint(hwnd, &ps);
            HBRUSH whiteBrush = CreateSolidBrush(RGB(255, 255, 255));
            FillRect(hdc, &ps.rcPaint, whiteBrush);
            DeleteObject(whiteBrush);
            // 3. Dibuixar el cercle
            HBRUSH brush = CreateSolidBrush(RGB(rand(255), rand(255), rand(255))); 
            SelectObject(hdc, brush);

            Ellipse(hdc, x - RADIUS, y - RADIUS, x + RADIUS, y + RADIUS);

            EndPaint(hwnd, &ps);
            DeleteObject(brush);

            return 0;
        }
        /*case WM_MOUSEMOVE: {
            x = LOWORD(lParam);
            y = HIWORD(lParam);

            InvalidateRect(hwnd, NULL, TRUE); /* This forces the invalidation of the window and, so, the WM_PAINT message is sent again*/
        //    return 0;
        //}
        case WM_DESTROY:
        {
            KillTimer(hwnd, IDT_TIMER1);
            PostQuitMessage(0);
            return 0;
            break;
        }
        default:
            return DefWindowProcA(hwnd, uMsg, wParam, lParam);
        }
    return DefWindowProcA(hwnd, uMsg, wParam, lParam);
}

/**
 * This is actually the entry point for this program.
 * This replaces the int main method with a more specific one for us in this case.
 */
int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrev, PSTR lpCmdLine, int nCmdShow)
{

    WNDCLASSA class = {
        0, /* Style */
        WinProc, /* A pointer to the window procedure */
        0, /* Extra bytes to allocate following the window-class structure*/
        0, /* Extra bytes to allocate to the window instance*/
        hInstance, /* A handle to the instance that contains the window procedure for the class */
        NULL, /* ICON, if NULL the system provides the default icon  */
        NULL, /* A handle to the class cursor */
        NULL, /* Hbrush, if null we are responsible for painting the background*/
        NULL, /* Resource name of the class menu */
        "FirstTest" /* Class name, identifier for the class */
    };

    RegisterClassA(&class);
    HWND windowHandle = CreateWindowA("FirstTest", "Hello World", WS_OVERLAPPEDWINDOW, CW_USEDEFAULT, CW_USEDEFAULT, WINDOW_WIDTH, WINDOW_HEIGHT, NULL, NULL, hInstance, NULL);
    BOOL showWindowRet = ShowWindow(windowHandle, nCmdShow);

    MSG msg; /* MSG structure that receives message information from the thread's message queue */
    for (;;) {
        /**
         * If we get the quit message we break from the loop
         */
        if (GetMessageA(&msg, NULL, 0, 0) == 0) break;
        else DispatchMessageA(&msg);
    }

    return 0;
}