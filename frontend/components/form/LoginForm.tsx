"use client";
import { useState } from "react";
import { useMutation } from "@apollo/client/react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { useAuth } from "@/providers/auth-provider";
import { LOGIN_MUTATION } from "@/lib/graphql/mutations";

export default function LoginForm() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState("");
  const router = useRouter();
  const { login: authLogin } = useAuth();

  const [login, { loading }] = useMutation(LOGIN_MUTATION, {
    onCompleted: (data) => {
      authLogin(data.login.accessToken);
      console.log("Login successful:", data);
      router.push("/");
    },
    onError: (error) => {
      setError(error.message);
    },
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    login({ variables: { input: { email, password } } });
  };

  return (
    <div className="w-full max-w-md p-8 space-y-6 bg-[#FFFFFF] rounded-lg shadow-md">
      <h1 className="text-2xl font-bold text-center text-[#1F1F1F]">Fazer login</h1>
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
            placeholder="E-mail"
            className="block w-full px-3 py-2 mt-1 text-[#595959] font-medium placeholder-[#858787] border border-[#858787] rounded-[4px] shadow-sm appearance-none focus:outline-none focus:ring-indigo-500 focus:border-[#0B57D0] focus:border-2 text-base"
          />
        </div>
        <div>
          <label
            htmlFor="password"
            className="block text-sm font-medium text-gray-700"
          >
            Password
          </label>
          <div className="relative">
            <input
              id="password"
              name="password"
              type={showPassword ? "text" : "password"}
              autoComplete="current-password"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="password"
              className="block w-full px-3 py-2 mt-1 text-[#595959] font-medium placeholder-[#858787] border border-[#858787] rounded-[4px] shadow-sm appearance-none focus:outline-none focus:ring-indigo-500 focus:border-[#0B57D0] focus:border-2 text-base"

            />
            <button
              type="button"
              onClick={() => setShowPassword(!showPassword)}
              className="absolute inset-y-0 right-0 px-3 text-black cursor-pointer flex items-center text-sm leading-5"
            >
              {showPassword ? "Hide" : "Show"}
            </button>
          </div>
        </div>
        {error && <p className="text-sm text-red-600">{error}</p>}
        <div>
          <button
            type="submit"
            disabled={loading}
            className="flex justify-center cursor-pointer w-full px-4 py-2 text-sm font-medium text-white bg-[#0B57D0] border border-transparent rounded-[50px] shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            {loading ? "Logging in..." : "Login"}
          </button>
        </div>
      </form>
      <div className="text-sm text-center">
        <Link href="/auth/register" className="font-medium text-[#0b57d0] rounded-[8px] p-1 hover:bg-gray-200">
          Don&apos;t have an account? Sign up
        </Link>
      </div>
      <div className="text-sm text-center">
        <Link href="/auth/forgot-password" className="font-medium text-[#0b57d0] rounded-[8px] p-1 hover:bg-gray-200">
          Forgot your password?
        </Link>
      </div>
    </div>
  );
}
