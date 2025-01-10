"use client";
import Link from "next/link";
import ForumIcon from "@mui/icons-material/Forum";
import { SvgIconComponent } from "@mui/icons-material";
import PostAddIcon from "@mui/icons-material/PostAdd";
import LogoutIcon from "@mui/icons-material/Logout";
import { usePathname } from "next/navigation";
type NavData = { href: string; label: string; icon: SvgIconComponent };
const navItemData: NavData[] = [
  { href: "/", label: "Forum", icon: ForumIcon },
  { href: "/post/new", label: "Create Post", icon: PostAddIcon },
  { href: "/logout", label: "Logout", icon: LogoutIcon },
];
export default function Navbar() {
  const pathname = usePathname();
  return (
    <div className="shadow-md bg-slate-50 min-w-48">
      <p className="text-xl font-semibold">CVWO Forum</p>

      {navItemData.map((val) => (
        <NavItem
          key={val.href}
          href={val.href}
          label={val.label}
          Icon={val.icon}
          selected={pathname == val.href}
        />
      ))}
    </div>
  );
}
function NavItem({
  href,
  label,
  Icon,
  selected,
}: {
  href: string;
  label: string;
  Icon: SvgIconComponent;
  selected: boolean;
}) {
  return (
    <Link href={href}>
      <div
        className={
          "flex gap-6 p-2 hover:bg-slate-300 " +
          (selected ? "bg-slate-600" : "")
        }
      >
        <Icon />
        <p>{label}</p>
      </div>
    </Link>
  );
}
