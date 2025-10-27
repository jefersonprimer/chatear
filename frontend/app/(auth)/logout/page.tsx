"use client";
import { useEffect } from "react";
import { useMutation } from "@apollo/client";
import { useRouter } from "next/navigation";
import { LOGOUT_MUTATION } from "@/lib/graphql/mutations";

export default function LogoutPage() {
  const router = useRouter();
  const [logout, { loading, error }] = useMutation(LOGOUT_MUTATION, {
    onCompleted: () => {
      // TODO: Clear the tokens from local storage or cookies
      router.push("/auth/login");
    },
    onError: (error) => {
      console.error("Logout failed:", error);
      // Even if logout fails, redirect to login
      router.push("/auth/login");
    },
  });

  useEffect(() => {
    logout();
  }, [logout]);

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 space-y-6 text-center bg-white rounded-lg shadow-md">
        <h1 className="text-2xl font-bold text-gray-900">Logging out</h1>
        {loading && <p>Please wait...</p>}
        {error && <p className="text-red-600">An error occurred while logging out.</p>}
      </div>
    </div>
  );
}