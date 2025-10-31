
import jwt from "jsonwebtoken";

const SECRET = process.env.JWT_SECRET;

if (!SECRET) {
  throw new Error("JWT_SECRET is not defined in your environment variables");
}

export async function verifyAuth(token: string) {
  try {
    return jwt.verify(token, SECRET);
  } catch {
    return null;
  }
}
