using System;
using System.Globalization;

namespace CourseApp.Module3
{
    public class CycleOfStr
    {
        public static int GetHash(string s, int n)
        {
            int res = 0, p = 2803, x = 31;
            for (int i = 0; i < n; i++)
            {
                res = ((res * x) + s[i]) % p;
            }

            return res;
        }

        public static int RabinKarp(string s, string t)
        {
            int ht = GetHash(t, t.Length);
            int hs = GetHash(s, s.Length);
            int p = 2803, x = 27;

            for (int i = 0; i < s.Length; i++)
            {
                if (ht == hs)
                {
                    return i;
                }

                ht = ((ht * x) - (t[i] * ht) + t[i]) % p;
                ht = (ht + p) % p;
            }

            return -1;
        }

        public static void Pain() // Main method
        {
            var s = Console.ReadLine();
            var t = Console.ReadLine();

            var res = RabinKarp(s, t);
            Console.WriteLine(res);
        }
    }
}