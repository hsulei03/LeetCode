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

            Solution_PalindromeNumber test = new Solution_PalindromeNumber();
            var result = test.IsPalindrome(12321);
            Console.WriteLine(result);
            Console.ReadLine();
            
        }
    }
}
