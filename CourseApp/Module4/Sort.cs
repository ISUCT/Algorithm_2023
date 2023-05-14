using System;
using System.Collections.Generic;
using System.Linq;

namespace CourseApp.Module4
{
    public class Sort
    {
        public static void Pain()
        {
            var count = Convert.ToInt16(Console.ReadLine());
            var mass = Console.ReadLine().Split(" ");

            var a = Pup(count, mass);

            var str = a ? "YES" : "NO";

            Console.WriteLine(str);
        }

        public static bool Pup(int count, string[] mass)
        {
            var res = false;
            var stack = new Stack<int>();
            var ind = 0;
            var start = 1;

            while (count >= ind + 1)
            {
                stack.Push(Convert.ToInt16(mass[ind]));
                ind++;

                while (stack.Count != 0 && stack.Peek() == start)
                {
                    stack.Pop();
                    start++;
                }
            }

            if (stack.Count == 0)
            {
                res = true;
            }

            return res;
        }
    }
}