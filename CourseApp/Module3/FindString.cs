using System;
using System.Collections.Generic;
using System.Linq;

namespace CourseApp.Module3
{
    public class FindString
    {
        public static void FindStr()
        {
            var a = Console.ReadLine();
            var b = Console.ReadLine();
            List<string> result = new List<string>();
            for (int i = 0; ; i += b.Length)
            {
                result.Add(a.IndexOf(b, i).ToString());
            }

            Console.WriteLine(result.Aggregate((d, c) => d + " " + c) + "\r\n");
        }
    }
}
