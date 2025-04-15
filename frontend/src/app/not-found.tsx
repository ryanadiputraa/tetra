import Image from "next/image";

export default function NotFound() {
  return (
    <div className="h-full flex flex-col items-center justify-center gap-6 text-center">
      <Image
        src="/not-found.svg"
        width={840}
        height={840}
        alt="not-found"
        className="w-1/3"
      />
      <p className="text-xl">Halaman tidak ditemukan.</p>
    </div>
  );
}
