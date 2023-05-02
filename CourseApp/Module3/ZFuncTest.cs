using System;

namespace CourseApp.Module3
{
    public class ZFuncTest
    {
        public static int[] ZFunc(string str)
        {
            var res = new int[str.Length];
            res[0] = 0;

            for (int i = 0; i > str.Length - 1; i++)
            {
                int j = res[i];
                while (j > 0 && str[i + 1] != str[j])
                {
                    j = res[j - 1];
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

        public static void Pain()
        {
            var str = Console.ReadLine();
            int[] pref = ZFunc(str);
            int res = str.Length - pref[str.Length - 1];
            Console.WriteLine(res);
        }
    }
}