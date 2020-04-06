using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeSolutions
{
    public class Solution_PalindromeNumber
    {
        ///Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
        ///Coud you solve it without converting the integer to a string?
        /// <summary>
        /// 轉string後做判斷
        /// </summary>
        /// <param name="x"></param>
        /// <returns></returns>
        public bool IsPalindrome_1(int x)
        {
            var target = x.ToString();
            for (int i = 0; i < target.Length/2; i++)
            {
                if(target[i] != target[target.Length-1 - i])
                {
                    return false;
                }                
            }
            return true;

        }

        /// <summary>
        /// 不轉string
        /// </summary>
        /// <param name="x"></param>
        /// <returns></returns>
        public bool IsPalindrome(int x)
        {
            int result = 0;
            if (x < 0)
            {
                return false;
            }
            else
            {
                int input = x;
                while (input > 0)
                {
                    result = (result * 10) + (input % 10);
                    input = input / 10;
                }
                return x == result ? true : false;
            }

        }
    }
}
