using System;
using System.Collections.Generic;

namespace CourseApp
{
    public class PSP
    {
        public static void Psp()
        {
            var str = Console.ReadLine();
            var stack = new Stack<char>();

            foreach (char symbol in str)
            {
                if (symbol == '(')
                {
                    stack.Push(symbol);
                }
                else if (symbol == ')')
                {
                    if (stack.Count > 0 && stack.Peek() == '(')
                    {
                        stack.Pop();
                    }
                    else
                    {
                        stack.Push(symbol);
                    }
                }
            }

            Console.WriteLine(stack.Count);
        }
    }
}