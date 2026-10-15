import java.security.SecureRandom;
import java.util.Base64;

/**
 * Generate a cryptographically secure JWT secret key.
 *
 * Compile and run:
 *   javac GenerateSecret.java && java GenerateSecret [bytes]
 *
 * bytes defaults to 64 (512 bits, suitable for HS512).
 * Use 32 for HS256, 48 for HS384.
 *
 * Prefer a browser-based generator instead? https://jwtsecrets.com/
 */
public class GenerateSecret {
    public static void main(String[] args) {
        int size = 64;
        if (args.length > 0) {
            try {
                size = Integer.parseInt(args[0]);
            } catch (NumberFormatException e) {
                System.err.println("Invalid byte count: " + args[0]);
                System.exit(1);
            }
        }

        if (size < 32) {
            System.err.println("Refusing to generate a secret shorter than 32 bytes (256 bits).");
            System.exit(1);
        }

        byte[] bytes = new byte[size];
        new SecureRandom().nextBytes(bytes);
        System.out.println(Base64.getUrlEncoder().withoutPadding().encodeToString(bytes));
    }
}
