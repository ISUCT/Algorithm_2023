using System;
using CourseApp.Module2;

namespace CourseApp
{
    public class Program
    {
        public static void Main()
        {
            var a = Inversion.Parse();
            Inversion.Sort(a);

            /*var massLen = Convert.ToInt32(Console.ReadLine());
            var mass = Console.ReadLine().Split(" ");
            var list = new int[massLen];
            for ( int i = 0; i < massLen; i++)
            {
                list[i] = Convert.ToInt32(mass[i]);
            }

            MergeBorderOutput.BMergeSort(list);
            foreach ( int i in MergeBorderOutput.BMergeSort(list))
            {
                Console.Write(i + " ");
            }*/
        }
    }
}