using System;
using Xunit;
using CourseApp;

namespace CourseApp.Tests
{
    public class DemoTest
    {
        [Fact]
        public void Test1()
        {
            Assert.Equal("22", CourseApp.Program.Plus(10, 12).ToString());
        }

        [Fact]
        public void Test2()
        {
            Assert.Equal("2", CourseApp.Program.Plus(1, 1).ToString());
        }

        [Fact]
        public void Test3()
        {
            Assert.Equal("20000", CourseApp.Program.Plus(10000, 10000).ToString());
        }
    }
}
