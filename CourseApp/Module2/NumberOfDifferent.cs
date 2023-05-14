using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace CourseApp.Module2
{
    public class NumberOfDifferent
    {
        public static int[] Parse()
        {
            var len = Convert.ToInt32(Console.ReadLine());
            var mass = new int[len];
            var str = Console.ReadLine().Split(' ');
            for (int i = 0; i < len; i++)
            {
                mass[i] = Convert.ToInt32(str[i]);
            }

            return mass;
        }

        public static void CountDiffrent(int[] mass)
        {
            Console.WriteLine(mass.ToList().Distinct().Count());
        }
    }
}
