using System;
using System.IO;
using Xunit;
using CourseApp.Module3;

namespace CourseApp.Tests.Module3
{
    [Collection("Sequential")]
    public class ZFuncTest1 : IDisposable
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
        [InlineData(
            "zbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzbzz",
            113)]
        [InlineData("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz", 1)]

        public void Test1(string input, int expected)
        {
            var stringWriter = new StringWriter();
            Console.SetOut(stringWriter);

            var stringReader = new StringReader(input);
            Console.SetIn(stringReader);

            ZFuncTest.Pain();

            var output = stringWriter.ToString().Split(Environment.NewLine, StringSplitOptions.RemoveEmptyEntries);
            Assert.Equal($"{expected}", output[0]);
            var standardOutput = new StreamWriter(Console.OpenStandardOutput());
        }
    }
}