"use client";
import { useState, useEffect } from "react";
import { useMutation } from "@apollo/client/react";
import { useRouter, useSearchParams } from "next/navigation";
import Link from "next/link";
import { VERIFY_EMAIL_MUTATION } from "@/lib/graphql/mutations";

export default function VerifyEmailPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [token, setToken] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  const [verifyEmail, { loading }] = useMutation(VERIFY_EMAIL_MUTATION, {
    onCompleted: () => {
      setSuccess(true);
      setTimeout(() => {
        router.push("/auth/login");
      }, 3000);
    },
    onError: (error) => {
      setError(error.message);
    },
  });

  useEffect(() => {
    const tokenFromUrl = searchParams.get("token");
    if (tokenFromUrl) {
      setToken(tokenFromUrl);
      verifyEmail({ variables: { input: { token: tokenFromUrl } } });
    } else {
      setError("Token not found in URL.");
    }
  }, [searchParams, verifyEmail]);

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md">
        <h1 className="text-2xl font-bold text-center text-gray-900">Verify Email</h1>
        {loading && <p className="text-center">Verifying your email...</p>}
        {error && <p className="text-center text-red-600">Error: {error}</p>}
        {success && (
          <p className="text-center text-green-600">
            Your email has been successfully verified. You will be redirected to the login page shortly.
          </p>
        )}
        <div className="text-sm text-center">
          <Link href="/auth/login" className="font-medium text-indigo-600 hover:text-indigo-500">
            Back to login
          </Link>
        </div>
      </div>
    </div>
  );
}
