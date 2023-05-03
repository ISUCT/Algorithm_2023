using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using CourseApp.Module2;
using CourseApp.Module3;
using CourseApp.Module4;

namespace CourseApp
{
    public class Program
    {
        public static void Main()
        {
            var i = BalsCount.Count("5 1 3 3 3 2");
            Console.WriteLine(i);
        }
    }
}