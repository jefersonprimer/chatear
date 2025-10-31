"use client";
import { useState } from "react";
import { useMutation } from "@apollo/client/react";
import Link from "next/link";
import { RESET_PASSWORD_MUTATION } from "@/lib/graphql/mutations";

export default function ForgotPasswordForm() {
  const [email, setEmail] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  const [resetPassword, { loading }] = useMutation(RESET_PASSWORD_MUTATION, {
    onCompleted: () => {
      setSuccess(true);
    },
    onError: (error) => {
      setError(error.message);
    },
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setSuccess(false);
    resetPassword({ variables: { input: { email } } });
  };

  return (
    <div className="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md">
      <h1 className="text-2xl font-bold text-center text-gray-900">Forgot Password</h1>
      {success ? (
        <p className="text-center text-green-600">
          If an account with that email exists, we have sent you an email with instructions to reset your password.
        </p>
      ) : (
        <form className="space-y-6" onSubmit={handleSubmit}>
          <div>
            <label
              htmlFor="email"
              className="block text-sm font-medium text-gray-700"
            >
              Email
            </label>
            <input
              id="email"
              name="email"
              type="email"
              autoComplete="email"
              required
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="block w-full px-3 py-2 mt-1 placeholder-gray-400 border border-gray-300 rounded-md shadow-sm appearance-none focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            />
          </div>
          {error && <p className="text-sm text-red-600">{error}</p>}
          <div>
            <button
              type="submit"
              disabled={loading}
              className="flex justify-center w-full px-4 py-2 text-sm font-medium text-white bg-indigo-600 border border-transparent rounded-md shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              {loading ? "Sending..." : "Send password reset email"}
            </button>
          </div>
        </form>
      )}
      <div className="text-sm text-center">
        <Link href="/auth/login" className="font-medium text-indigo-600 hover:text-indigo-500">
          Back to login
        </Link>
      </div>
    </div>
  );
}
