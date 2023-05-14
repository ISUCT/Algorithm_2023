using System;
using System.Collections.Generic;

public class Pairs
{
    public int Index { get; set; }

    public int Cost { get; set; }

    public static void PairSort()
    {
        var list = new List<Pairs>();
        int count = Convert.ToInt32(Console.ReadLine());
        int i = 0;
        while (i < count)
        {
            var str1 = Console.ReadLine();
            var index = Convert.ToInt32(str1.Split(" ")[0]);
            var cost = Convert.ToInt32(str1.Split(" ")[1]);
            var pair = new Pairs();
            pair.Index = index;
            pair.Cost = cost;
            list.Add(pair);
            i++;
        }

        for (int k = 0; k < count - 1; k++)
        {
            for (int j = 0; j < count - 1; j++)
            {
                if (list[j].Cost < list[j + 1].Cost)
                {
                    (list[j], list[j + 1]) = (list[j + 1], list[j]);
                }
            }
        }

       for (int k = 0; k < count - 1; k++)
       {
            for (int j = 0; j < count - 1; j++)
            {
                if ((list[j].Cost == list[j + 1].Cost) && (list[j].Index > list[j + 1].Index))
                {
                    (list[j], list[j + 1]) = (list[j + 1], list[j]);
                }
            }
       }

        foreach (var ind in list)
        {
            Console.WriteLine($"{ind.Index} {ind.Cost}");
        }
    }
}