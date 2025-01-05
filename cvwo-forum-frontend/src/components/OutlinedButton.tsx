import React from "react";

export default function OutlinedButton({
  type,
  disabled,
  children,
  onClick = () => {},
}: {
  type: "submit" | "reset" | "button" | undefined;
  disabled: boolean;
  children: React.ReactNode;
  onClick?: () => void;
}) {
  return (
    <button
      type={type}
      disabled={disabled}
      onClick={onClick}
      className="py-2 mt-4 border border-blue-500 rounded-lg hover:bg-blue-200 disabled:bg-gray-400"
    >
      {children}
    </button>
  );
}
