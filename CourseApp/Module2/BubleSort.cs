using System;
using System.Collections.Generic;
using System.Text;

namespace CourseApp.Module2
{
    public class BubleSort
    {
        public static void SortProc()
        {
            var k = 0;
            var massLen = Convert.ToInt32(Console.ReadLine());
            var mass = Console.ReadLine().Split(" ");
            foreach (string i in mass)
            {
                for (int index = 0; index < massLen - 1; index++)
                {
                    if (Convert.ToInt32(mass[index]) > Convert.ToInt32(mass[index + 1]))
                    {
                        var temp = mass[index];
                        mass[index] = mass[index + 1];
                        mass[index + 1] = temp;
                        foreach (string elem in mass)
                        {
                            Console.Write(elem + " ");
                            k++;
                        }

                        Console.WriteLine();
                    }
                }
            }

            if (k == 0)
            {
                Console.WriteLine(0);
            }
        }
    }
}
