using LeetCodeSolutions;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode
{
    class Program
    {
        static void Main(string[] args)
        {
            int[] nums = { 3, 3 };
            var target = 6;
            Solution_TwoSum tryrun = new Solution_TwoSum();
            int[] result = tryrun.TwoSum(nums, target);
            foreach( var i in result)
            {
                Console.WriteLine(i);
            }
            Console.ReadLine();
        }
    }
}
