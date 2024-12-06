using System.Security.Cryptography;
using System.Text;

namespace api.Services;

public class SecurityService
{
	public string ComputeHash(string str)
	{
		var sha256 = SHA256.Create();
		byte[] hash = sha256.ComputeHash(Encoding.UTF8.GetBytes(str));

		var sBuilder = new StringBuilder();

		for (int i = 0; i < hash.Length; i++)
		{
			sBuilder.Append(hash[i].ToString("x2"));
		}

		return sBuilder.ToString();
	}

    public static string Base64Encode(string plainText)
    {
        var plainTextBytes = System.Text.Encoding.UTF8.GetBytes(plainText);
        return System.Convert.ToBase64String(plainTextBytes);
    }
    public static string Base64Decode(string base64EncodedData)
    {
        var base64EncodedBytes = System.Convert.FromBase64String(base64EncodedData);
        return System.Text.Encoding.UTF8.GetString(base64EncodedBytes);
    }
}
