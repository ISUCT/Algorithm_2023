using System;
using System.Collections.Generic;

namespace CourseApp
{
    public class PSP
    {
        public static void Psp()
        {
            var s = Console.ReadLine();
            Stack<char> stack = new Stack<char>();
            foreach (char c in s)
            {
                if (c == '(')
                {
                    stack.Push(c);
                }
                else if (c == ')')
                {
                    if (stack.Count > 0 && stack.Peek() == '(')
                    {
                        stack.Pop();
                    }
                    else
                    {
                        stack.Push(c);
                    }
                }
            }

            Console.WriteLine(stack.Count);
        }
    }
}