using System;
using System.IO;
using Xunit;
using CourseApp.Module3;

namespace CourseApp.Tests.Module3
{
    [Collection("Sequential")]
    public class FindStringTest : IDisposable
    {
        public void Dispose()
        {
            var standardOut = new StreamWriter(Console.OpenStandardOutput());
            standardOut.AutoFlush = true;
            var standardIn = new StreamReader(Console.OpenStandardInput());
            Console.SetOut(standardOut);
            Console.SetIn(standardIn);
        }

        [Theory]
        [InlineData("ababbababa\naba", "0 5 7\r\n")]
        public void Test1(string input, string expected)
        {
            var stringWriter = new StringWriter();
            Console.SetOut(stringWriter);

            var stringReader = new StringReader(input);
            Console.SetIn(stringReader);

            // act
            FindString.FindStr();

            // assert
            var output = stringWriter.ToString().Split(Environment.NewLine, StringSplitOptions.RemoveEmptyEntries);
            Assert.Equal($"{expected}", output[0]);
            var standardOutput = new StreamWriter(Console.OpenStandardOutput());
        }
    }
}
