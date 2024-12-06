using System.Security.Cryptography;
using System.Text;

namespace api.Services;

public class HashPasswordService
{
    private string Hash(string password)
    {
        var bytePassword = Encoding.UTF8.GetBytes(password);
        using (var alg = SHA512.Create())
        {
            string hex = "";

            var hashValue = alg.ComputeHash(bytePassword);
            foreach (byte x in hashValue)
            {
                hex += String.Format("{0:x2}", x);
            }
            return hex;
        }
    }

    public string Make(string password)
    {
        return Hash(password);
    }

    public bool Compare(string hashedPassword, string providedPassword)
    {
        return hashedPassword == Hash(providedPassword);
    }
}