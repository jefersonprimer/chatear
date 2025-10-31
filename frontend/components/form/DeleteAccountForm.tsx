"use client";
import { useState } from "react";
import { useMutation } from "@apollo/client/react";
import { useRouter } from "next/navigation";
import { DELETE_ACCOUNT_MUTATION } from "@/lib/graphql/mutations";

export default function DeleteAccountForm() {
  const router = useRouter();
  const [error, setError] = useState("");

  // TODO: Get the user ID from the auth context
  const userId = "mock-user-id";

  const [deleteAccount, { loading }] = useMutation(DELETE_ACCOUNT_MUTATION, {
    onCompleted: () => {
      // TODO: Clear the tokens from local storage or cookies
      router.push("/auth/register");
    },
    onError: (error) => {
      setError(error.message);
    },
  });

  const handleDelete = () => {
    if (window.confirm("Are you sure you want to delete your account? This action cannot be undone.")) {
      deleteAccount({ variables: { input: { userID: userId } } });
    }
  };

  return (
    <div className="w-full max-w-md p-8 space-y-6 text-center bg-white rounded-lg shadow-md">
      <h1 className="text-2xl font-bold text-gray-900">Delete Account</h1>
      <p>Are you sure you want to delete your account? This action is irreversible.</p>
      {error && <p className="text-red-600">{error}</p>}
      <div>
        <button
          onClick={handleDelete}
          disabled={loading}
          className="flex justify-center w-full px-4 py-2 text-sm font-medium text-white bg-red-600 border border-transparent rounded-md shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
        >
          {loading ? "Deleting..." : "Delete My Account"}
        </button>
      </div>
    </div>
  );
}
