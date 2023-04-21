using System;
using System.Linq;

namespace CourseApp.Module3
{
    public class ShiftString
    {
        public static int Rabin_Karp(string str, string subStr)
        {
            if (str == subStr)
            {
                return 0; // случай, когда строки изначально совпали
            }

            subStr = string.Concat(Enumerable.Repeat(subStr, 2)); // умножаем подстроку (паттерн) на 2 (как в Python)

            long w = 13;
            long b = 1;
            long t = 100000000;

            long first_hash = 0;
            long second_hash = 0;
            long xt = 1;

            foreach (char i in str.Reverse())
            {
                first_hash = (first_hash + (i * b)) % t; // высчитываем хэш 1 строки
                b = (b * w) % t;
            }

            b = 1;

            for (int i = str.Length - 1; i >= 0; i--)
            {
                second_hash = (second_hash + (subStr[i] * b)) % t; // высчитываем хэш 2 строки
                b = (b * w) % t;
            }

            for (int i = 0; i < str.Length - 1; i++)
            {
                xt = (xt * w) % t;
            }

            for (int i = 1; i < subStr.Length - str.Length + 1; i++)
            {
                if (first_hash == second_hash)
                {
                    return i - 1; // находим циклическое
                }

                second_hash = w * (second_hash - (subStr[i - 1] * xt)); // второй хэш
                second_hash += subStr[i + str.Length - 1];
                second_hash = second_hash % t; // избавляемся от переполнения

                if ((second_hash < 0 && t > 0) || (second_hash > 0 && t < 0))
                {
                    second_hash += t;
                }
            }

            return -1; // ничего не нашли(((
        }

        public static void EnterValues()
        {
            string r = Console.ReadLine();
            string q = Console.ReadLine();

            int result = Rabin_Karp(r, q);

            Console.WriteLine(result);
        }
    }
}