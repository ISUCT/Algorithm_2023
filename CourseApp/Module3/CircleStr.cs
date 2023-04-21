using System;
using System.Linq;

namespace CourseApp.Module3
{
    public class CircleStr
    {
        public static int Pain() // Main
        {
            var res = 0;
            var str = Console.ReadLine();
            var pup = KnutMP(str);
            if (str.Distinct().Count() == 1)
            {
                res = str.Length;
            }
            else
            {
                res = str.Length - pup[str.Length - 1];
            }

            Console.WriteLine(res);
            return res;
        }

        public static int[] KnutMP(string str)
        {
            int j = -1;
            var res = new int[str.Length];
            for (int i = 0; i < str.Length - 1; i++)
            {
                j = res[i];

                while (j > 0 && str[i + 1] != str[j])
                {
                    j = res[i - 1];
                }

                if (str[i + 1] == str[j])
                {
                    res[i + 1] = j + 1;
                }
                else
                {
                    res[i + 1] = 0;
                }
            }

            return res;
        }
    }
}
