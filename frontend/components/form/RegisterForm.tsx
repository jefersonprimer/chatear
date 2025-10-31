"use client";
import { useState, useRef } from "react";
import { useMutation } from "@apollo/client/react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import Image from "next/image";
import { REGISTER_USER } from "@/lib/graphql/mutations/registerUser";
import ReactCrop, { type Crop, centerCrop, makeAspectCrop } from "react-image-crop";
import "react-image-crop/dist/ReactCrop.css";
import camera from "@/public/camera.png";

export default function RegisterForm() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [gender, setGender] = useState("");
  const [error, setError] = useState("");
  const [avatar, setAvatar] = useState<File | null>(null);
  const [avatarPreview, setAvatarPreview] = useState<string>(camera.src);
  const [crop, setCrop] = useState<Crop>();
  const [showCropModal, setShowCropModal] = useState(false);
  const [termsAccepted, setTermsAccepted] = useState(false);
  const [uploadLoading, setUploadLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const imgRef = useRef<HTMLImageElement>(null);

  function onImageLoad(e: React.SyntheticEvent<HTMLImageElement>) {
    const { width, height } = e.currentTarget;
    const crop = centerCrop(
      makeAspectCrop(
        {
          unit: 'px',
          width: Math.min(width, height, 200),
        },
        1,
        width,
        height,
      ),
      width,
      height,
    );
    setCrop(crop);
  }

  const [registerUser, { loading: registerLoading }] = useMutation(
    REGISTER_USER,
    {
      onCompleted: () => {
        setSuccess(true);
      },
      onError: (error) => {
        setError(error.message);
      },
    }
  );



  const handleAvatarChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      setAvatar(e.target.files[0]);
      setAvatarPreview(URL.createObjectURL(e.target.files[0]));
      setShowCropModal(true);
    }
  };

  const getCroppedImg = (): Promise<File> => {
    return new Promise((resolve, reject) => {
      const image = imgRef.current;
      if (!image || !crop) {
        return reject(new Error("Image or crop not available"));
      }
      const canvas = document.createElement("canvas");
      const scaleX = image.naturalWidth / image.width;
      const scaleY = image.naturalHeight / image.height;
      canvas.width = crop.width;
      canvas.height = crop.height;
      const ctx = canvas.getContext("2d");

      if (!ctx) {
        return reject(new Error("Canvas context not available"));
      }

      ctx.drawImage(
        image,
        crop.x * scaleX,
        crop.y * scaleY,
        crop.width * scaleX,
        crop.height * scaleY,
        0,
        0,
        crop.width,
        crop.height
      );

      canvas.toBlob(
        (blob) => {
          if (!blob) {
            return reject(new Error("Canvas is empty"));
          }
          const file = new File([blob], "avatar.jpg", { type: "image/jpeg" });
          resolve(file);
        },
        "image/jpeg",
        1
      );
    });
  };

  const handleCropComplete = async () => {
    try {
      const croppedImage = await getCroppedImg();
      setAvatar(croppedImage);
      setAvatarPreview(URL.createObjectURL(croppedImage));
      setShowCropModal(false);
    } catch (error) {
      console.error("Error cropping image:", error);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    try {
      const { data: registerData } = await registerUser({
        variables: { input: { name, email, password, gender } },
      });

      if (registerData.registerUser && avatar) {
        setUploadLoading(true);
        const formData = new FormData();
        formData.append("operations", JSON.stringify({
          query: "mutation UploadAvatar($file: Upload!) { uploadAvatar(file: $file) }",
          variables: { file: null },
        }));
        formData.append("map", JSON.stringify({ "0": ["variables.file"] }));
        formData.append("0", avatar);

        await fetch(process.env.NEXT_PUBLIC_GRAPHQL_URL as string, {
          method: "POST",
          headers: {
            Authorization: `Bearer ${registerData.registerUser.accessToken}`,
          },
          body: formData,
        });
        setUploadLoading(false);
      }

    } catch (error: any) {
      setUploadLoading(false);
      setError(error.message);
    }
  };

  const isFormValid = name && email && password && gender && termsAccepted;

  return (
    <div className="w-full max-w-md p-8 space-y-6 bg-[#FFFFFF] rounded-lg shadow-md">
      <h1 className="text-2xl font-bold text-center text-[#1F1F1F]">Criar uma conta no WeChat</h1>
      {success ? (
        <p className="text-center text-green-600">
          We've sent a verification link to your email address. Please check your inbox and click the link to complete your registration.
        </p>
      ) : (
      <form className="space-y-6" onSubmit={handleSubmit}>
        <div className="flex justify-center">
          <label htmlFor="avatar-upload" className="cursor-pointer">
            <Image
              src={avatarPreview}
              alt="Avatar"
              width={100}
              height={100}
            />
          </label>
          <input
            id="avatar-upload"
            type="file"
            accept="image/*"
            onChange={handleAvatarChange}
            className="hidden"
          />
        </div>
        {showCropModal && (
          <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
            <div className="p-4 bg-white rounded-lg">
              <ReactCrop crop={crop} onChange={c => setCrop(c)}>
                <img src={avatarPreview} ref={imgRef} alt="Crop preview" onLoad={onImageLoad} />
              </ReactCrop>
              <div className="flex justify-end mt-4 space-x-2">
                <button
                  type="button"
                  onClick={() => setShowCropModal(false)}
                  className="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-200 rounded-md hover:bg-gray-300"
                >
                  Cancel
                </button>
                <button
                  type="button"
                  onClick={handleCropComplete}
                  className="px-4 py-2 text-sm font-medium text-white bg-indigo-600 rounded-md hover:bg-indigo-700"
                >
                  OK
                </button>
              </div>
            </div>
          </div>
        )}
        <div>
          <label
            htmlFor="name"
            className="block text-sm font-medium text-gray-700"
          >
            Name
          </label>
          <input
            id="name"
            name="name"
            type="text"
            autoComplete="name"
            required
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="block w-full px-3 py-2 mt-1 text-[#595959] font-medium placeholder-[#858787] border border-[#858787] rounded-[4px] shadow-sm appearance-none focus:outline-none focus:ring-indigo-500 focus:border-[#0B57D0] focus:border-2 text-base"
          />
        </div>
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
        <div>
          <label
            htmlFor="gender"
            className="block text-sm font-medium text-gray-700"
          >
            Gender
          </label>
          <select
            id="gender"
            name="gender"
            required
            value={gender}
            onChange={(e) => setGender(e.target.value)}
            className="block w-full px-3 py-2 mt-1 text-[#595959] placeholder-gray-400 border border-gray-300 rounded-md shadow-sm appearance-none focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          >
            <option value="">Select Gender</option>
            <option value="MALE">Male</option>
            <option value="FEMALE">Female</option>
          </select>
        </div>
        <div className="flex items-center">
          <input
            id="terms"
            name="terms"
            type="checkbox"
            checked={termsAccepted}
            onChange={(e) => setTermsAccepted(e.target.checked)}
            className="w-4 h-4 text-indigo-600 border-gray-300 rounded focus:ring-indigo-500"
          />
          <label
            htmlFor="terms"
            className="block ml-2 text-sm text-gray-900"
          >
            I have read and accept the{" "}
            <Link href="/policies/terms-of-use" className="text-indigo-600 hover:text-indigo-500">
              legal terms
            </Link>
            {" "}
            <Link href="/policies/privacy-policy" className="text-indigo-600 hover:text-indigo-500">
              Politica de privacidade
            </Link>

          </label>
        </div>
        {error && <p className="text-sm text-red-600">{error}</p>}
        <div>
          <button
            type="submit"
            disabled={!isFormValid || registerLoading || uploadLoading}
            className="flex justify-center w-full px-4 py-2 text-sm font-medium text-white bg-indigo-600 border border-transparent rounded-md shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50"
          >
            {registerLoading || uploadLoading ? "Signing up..." : "Sign up"}
          </button>
        </div>
      </form>
      )}
      <div className="text-sm text-center">
        <Link href="/auth/login" className="font-medium text-indigo-600 hover:text-indigo-500">
          Already have an account? Log in
        </Link>
      </div>
    </div>
  );
}
