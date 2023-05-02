using System;

namespace CourseApp.Module3
{
    public class CyclingString
    {
        public static int[] Prefix_function(string str) // Z-function finder
        {
            int[] res = new int[str.Length];
            res[0] = 0;

            for (int i = 0; i < str.Length - 1; i++)
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
            string str = Console.ReadLine();

            int[] prefixs = Prefix_function(str);

            int result = str.Length - prefixs[str.Length - 1];

            Console.WriteLine(result);
        }
    }
}