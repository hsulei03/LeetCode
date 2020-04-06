using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCodeSolutions
{
    /// <summary>
    /// Given a 32-bit signed integer, reverse digits of an integer.
    /// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^-31]
    /// For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
    /// </summary>
    public class Solution_ReverseInteger
    {
        public int Reverse(int x)
        {
            int result = 0;
            checked
            {
                try
                {
                    int input = Math.Abs(x);
                    while (input > 0)
                    {
                        result = (result * 10) + (input % 10);
                        input = input / 10;
                    }
                }
                catch (OverflowException)
                {
                    return 0;
                }
            }

            return x > 0 ? result : (0 - 1) * result;

        }
    }
}
