namespace CourseApp.Module5
{
    public class Tree
    {
        public Tree(int value)
        {
            this.Value = value;
            this.Left = null;
            this.Right = null;
        }

        public Tree Left { get;  set; }

        public Tree Right { get; set; }

        public int Value { get; set; }

        public int GetVal(int value)
        {
            return this.Value;
        }

        public void SetVal(int val)
        {
            this.Value = val;
        }

        public Tree GetLeft(Tree elem) => elem.Left == null ? null : elem.Left;
    }
}