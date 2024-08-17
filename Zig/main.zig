const std = @import("std");

const num1: u32 = 123;
const num2: u32 = 321;

const expectedSum: u32 = num1+num2;

pub fn main() void {
    const stdout = std.io.getStdOut().writer();
    const message = "Hello BounZig World!";
    stdout.print("{s}\n", .{message}) catch unreachable;
    stdout.print("Sum for {d} and {d} has returned {d} and was expected to return {d}\n", .{num1,num2,sumOFNumbers(num1, num2), expectedSum}) catch unreachable;
}

pub fn sumOFNumbers(num1ToSum: u32, num2ToSum: u32) u32 {
    return num1ToSum + num2ToSum;
}