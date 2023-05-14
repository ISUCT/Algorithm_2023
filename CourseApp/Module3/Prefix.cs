using System;
using System.Collections.Generic;
using System.Globalization;
using System.Reflection.PortableExecutable;
using System.Text;

namespace CourseApp.Module3
{
    public class Prefix
    {
        public static int[] Period(string str)
        {
            var length = str.Length;
            var result = new int[length];
            for (int i = 0; i < length - 1; i++)
            {
                var j = result[i];
                while ((j > 0) && (str[i + 1] != str[j]))
                {
                    j = result[i - 1];
                }

                if (str[i + 1] == str[j])
                {
                    result[i + 1] = j + 1;
                }
                else
                {
                    result[i + 1] = 0;
                }
            }

            return result;
        }

        public static void Pain()
        {
            var str = Console.ReadLine();
            var pref = Period(str);
            var res = str.Length - pref[str.Length - 1];

            if (str.Length % res == 0)
            {
                Console.WriteLine(str.Length / res);
            }
            else
            {
                Console.WriteLine(1);
            }
        }
    }
}
