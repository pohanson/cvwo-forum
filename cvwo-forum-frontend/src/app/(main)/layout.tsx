import Navbar from "@/components/navigation/Navbar";

export default function NavbarLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <div className="flex h-screen bg-slate-200">
      <Navbar />
      {children}
    </div>
  );
}
