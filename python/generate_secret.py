#!/usr/bin/env python3
"""Generate a cryptographically secure JWT secret key.

Usage:
    python3 generate_secret.py [bytes]

bytes defaults to 64 (512 bits, suitable for HS512).
Use 32 for HS256, 48 for HS384.

Prefer a browser-based generator instead? https://jwtsecrets.com/
"""
import secrets
import sys


def main() -> None:
    raw = sys.argv[1] if len(sys.argv) > 1 else "64"
    try:
        size = int(raw)
    except ValueError:
        print(f"Invalid byte count: {raw}", file=sys.stderr)
        raise SystemExit(1)

    if size < 32:
        print("Refusing to generate a secret shorter than 32 bytes (256 bits).", file=sys.stderr)
        raise SystemExit(1)

    # secrets.token_urlsafe uses the OS CSPRNG and returns a URL-safe string.
    print(secrets.token_urlsafe(size))


if __name__ == "__main__":
    main()
