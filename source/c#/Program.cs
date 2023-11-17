using C = System.Console;
using O = System.ConsoleColor;
using W = System.Diagnostics.Stopwatch;
using F = System.IO.File;
using L = System.Collections.Generic.HashSet<string>;
using R = System.Text.RegularExpressions.Regex;
using S = System.StringSplitOptions;

/// <summary>Класс программы</summary>
internal partial class Program
{

	/// <summary></summary>
	/// <returns></returns>
	[System.Text.RegularExpressions.GeneratedRegex(@"^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$")]
	private static partial R M();

	/// <summary>Точка входа</summary>
	/// <param name="a">Аргументы командной строки</param>
	private static void Main(string[] a)
	{
		W t = new();
		t.Start();
		C.Title = "Email Duplicator";
		C.CursorVisible = false;
		C.ForegroundColor = O.DarkGreen;
		C.WriteLine("*************************************");
		C.WriteLine("* Email Duplicator v. 2023-11-17    *");
		C.WriteLine("* https://github.com/Yxine/EmailDup *");
		C.WriteLine("*************************************");
		C.WriteLine();
		C.ForegroundColor = O.White;
		if (a.Length < 1)
		{
			t.Stop();
			C.WriteLine("  Command-line: EmailDup.exe file");
			C.ForegroundColor = O.DarkGray;
			C.WriteLine();
			C.WriteLine("  Press a key...");
			C.ReadKey();
			return;
		}
		var c = M();
		bool r(string e) => c.IsMatch(e);
		var f = a[0];
		var l = new L();
		var s = F.ReadAllLines(f);
		var b = new[] { ",", ";", " " };
		foreach (var e in s)
		{
			var m = e.Split(b, S.RemoveEmptyEntries);
			foreach (var i in m)
			{
				var w = i.ToLower();
				if (!l.Contains(w) && r(w)) l.Add(w);
			}
		}
		F.WriteAllLines($"{f}.cleared.txt", l);
		C.WriteLine($"  Start lines: {s.Length}");
		C.WriteLine($"  Final lines: {l.Count}");
		C.WriteLine();
		C.WriteLine($"  Time taken:  {t.Elapsed}");
		t.Stop();
	}

}
