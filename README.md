# JWT Secret Toolkit

Generate cryptographically secure JWT secret keys and signing material — in Node.js, Python, Java, or Go. One tiny, dependency-light snippet per language, all backed by each platform's CSPRNG.

> **Why this exists:** the most common JWT vulnerability in production is a weak or leaked signing secret. This repo gives you the correct, copy-paste way to generate a strong one in the language you already use.
>
> Need to generate a key right now, without installing anything? Use the free browser-based **[JWT Secret Key Generator](https://jwtsecrets.com/)** — it runs fully client-side and never transmits your key.

📖 **Read the full guide:** [How to Generate a Secure JWT Secret Key in 2026 (Node.js, Python, Java)](https://dev.to/chubiao_huangmrhuang_/how-to-generate-a-secure-jwt-secret-key-in-2026-nodejs-python-java-5a7d) — covers key length per algorithm, using the secret in `jsonwebtoken` and PyJWT, the three mistakes that leak a secret, and symmetric vs. asymmetric key choice.

## What you get

| Language | File | CSPRNG used |
|----------|------|-------------|
| Node.js  | [`node/generate-secret.js`](node/generate-secret.js) | `crypto.randomBytes` |
| Python   | [`python/generate_secret.py`](python/generate_secret.py) | `secrets.token_hex` |
| Java     | [`java/GenerateSecret.java`](java/GenerateSecret.java) | `java.security.SecureRandom` |
| Go       | [`go/main.go`](go/main.go) | `crypto/rand` |

## Quick start

### Node.js

```bash
node node/generate-secret.js          # default: 64 bytes (512 bits)
node node/generate-secret.js 32       # 32 bytes (256 bits, for HS256)
```

### Python

```bash
python3 python/generate_secret.py     # default: 64 bytes
python3 python/generate_secret.py 32  # 32 bytes
```

### Go

```bash
cd go && go run main.go               # prints a 64-byte secret
```

### Java

```bash
cd java && javac GenerateSecret.java && java GenerateSecret
```

## How long should a JWT secret be?

Match the key to your signing algorithm:

| Algorithm | Minimum | This toolkit's default |
|-----------|---------|------------------------|
| HS256     | 256 bits (32 bytes) | — |
| HS384     | 384 bits (48 bytes) | — |
| HS512     | 512 bits (64 bytes) | **64 bytes** |

Never use `Math.random()` (JavaScript), `random` (Python), or a human-memorable passphrase as a JWT secret. Always use the platform CSPRNG, as this toolkit does.

## Handling the generated secret

1. **Store it outside source control** — environment variable or a secrets manager (AWS Secrets Manager, Google Secret Manager, HashiCorp Vault).
2. **Use a different secret per environment** (dev / staging / prod) and per service.
3. **Rotate on a schedule**, and immediately after any suspected exposure.

```bash
# Example: write to a .env file you have already gitignored
node node/generate-secret.js > /tmp/secret.txt
echo "JWT_SECRET=$(cat /tmp/secret.txt)" >> .env
```

## Browser alternative

If you just need a key and do not want to run anything locally, the **[jwtsecrets.com JWT Secret Key Generator](https://jwtsecrets.com/)** produces the same quality of key entirely in your browser. Related tools: [JWT Decoder](https://decoder.jwtsecrets.com/), [JWT Encoder](https://jwtsecrets.com/tools/jwt-encode).

## License

MIT — see [LICENSE](LICENSE).
