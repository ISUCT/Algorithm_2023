using System;

namespace CourseApp.Module3
{
    public class CycleOfStr
    {
        public static int[] Kmp(string str)
        {
            int[] p = new int[str.Length];
            for (var i = 1; i < str.Length; i++)
            {
                int j = p[i - 1];
                while (j > 0 && str[i] != str[j])
                {
                    j = p[j - 1];
                }

                if (str[i] == str[j])
                {
                    j++;
                }

                p[i] = j;
            }

            return p;
        }

        public static int Find(string s, string t)
        {
            int n = s.Length;
            string r = t + "#" + s + s;
            int[] p = Kmp(r);
            if (p[n + 1 + t.Length - 1] == t.Length)
            {
                return n - t.Length;
            }
            else
            {
                return -1;
            }
        }

        public void Ass()
        {
            var str = Console.ReadLine();
            var subStr = Console.ReadLine();
            var res = Find(str, subStr);
            Console.WriteLine(res);
        }
    }
}