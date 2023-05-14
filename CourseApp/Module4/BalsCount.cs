using System;
using System.Collections.Generic;

namespace CourseApp.Module4
{
    public class BalsCount
    {
        public static void Count()
        {
            var nums = Console.ReadLine();
            var myStack = new Stack<int>();
            var resStack = new Stack<int>();
            var count = 1;
            var res = 0;

            for (int i = 2; i < nums.Length; i += 2)
            {
                myStack.Push(Convert.ToInt16(nums[i]));
            }

            var length = myStack.Count;

            while (myStack.Count > 0)
            {
                resStack.Push(myStack.Peek());
                if (resStack.Pop() == myStack.Pop())
                {
                    count++;
                    if (count == 3)
                    {
                        res++;
                        resStack.Clear();
                    }
                }
                else
                {
                    resStack.Clear();
                    count = 1;
                }
            }

            Console.WriteLine(res * 3);
        }
    }
}
