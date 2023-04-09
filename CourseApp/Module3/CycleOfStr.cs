using System;

namespace CourseApp.Module3
{
    public class CycleOfStr
    {
        public static int[] Kmp(string s)
        {
            int n = s.Length;
            int[] pi = new int[n];
            for (int i = 1; i < n; i++)
            {
                int j = pi[i - 1];
                while (j > 0 && s[i] != s[j])
                {
                    j = pi[j - 1];
                }

                if (s[i] == s[j])
                {
                    j++;
                }

                pi[i] = j;
            }

            return pi;
        }

        public static int Find(string s, string t)
        {
            int n = s.Length;
            int m = t.Length;
            if (n != m)
            {
                return -1;
            }

            s = s + '#' + s;
            int[] pi = Kmp(s);
            for (int i = 0; i < m; i++)
            {
                int j = pi[n + i + 1];
                if (t[i] != s[j])
                {
                    return -1;
                }
            }

            return n - pi[n * 2];
        }
    }
}