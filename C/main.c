#include <Windows.h>

/**
 * Extracted from https://www.youtube.com/watch?v=RiweaH6Qmro
 */

LRESULT WinProc(HWND hwnd, UINT uMsg, WPARAM wParam, LPARAM lParam)
{
    switch (uMsg) {
        case WM_PAINT:
            HDC DeviceContext = GetDC(hwnd);
            RECT rect = {
                75, /* Initial coordinates X */
                75, /* Initial coordinates Y */
                250, /* Final coordinates X */
                250, /* Final coordinates Y*/
            };
            HBRUSH solidBrush = CreateSolidBrush(RGB(63,139,139));
            FillRect(DeviceContext, &rect, solidBrush);
            ReleaseDC(hwnd, DeviceContext); /* Garbage collector for the DC. SUPER IMPORTANT*/
            DeleteObject(solidBrush); /* This function deletes a logical pen, brush, font or others, freeing some system resources*/
            return 0;
            break;
        case WM_DESTROY:
            PostQuitMessage(0);
            return 0;
            break;
        default:
            return DefWindowProcA(hwnd, uMsg, wParam, lParam);
    }
    return DefWindowProcA(hwnd,uMsg, wParam, lParam);
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
    HWND windowHandle = CreateWindowA("FirstTest", "Hello World", WS_CAPTION | WS_POPUP | WS_SYSMENU, 50, 50, 900, 600, NULL, NULL, hInstance, NULL );
    BOOL showWindowRet = ShowWindow(windowHandle, nCmdShow);

    MSG msg; /* MSG structure that receives message information from the thread's message queue */
    for(;;) {
        /**
         * If we get the quit message we break from the loop
         */
        if(GetMessageA(&msg, NULL, 0, 0) == 0) break;
        else DispatchMessageA(&msg);
    }

    return 0;
}