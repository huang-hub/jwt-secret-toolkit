#!/usr/bin/env node
/**
 * Generate a cryptographically secure JWT secret key.
 *
 * Usage:
 *   node generate-secret.js [bytes]
 *
 * bytes defaults to 64 (512 bits, suitable for HS512).
 * Use 32 for HS256, 48 for HS384.
 *
 * Prefer a browser-based generator instead? https://jwtsecrets.com/
 */
const crypto = require('crypto');

const bytes = Number.parseInt(process.argv[2], 10) || 64;

if (bytes < 32) {
  console.error('Refusing to generate a secret shorter than 32 bytes (256 bits).');
  process.exit(1);
}

const secret = crypto.randomBytes(bytes).toString('base64url');
console.log(secret);
