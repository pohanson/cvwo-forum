export default function TextInput({
  id,
  label,
  defaultValue,
}: {
  id: string;
  label: string;
  defaultValue?: string;
}) {
  return (
    <div className="flex flex-col gap-1 my-4">
      <label htmlFor={id}>{label}</label>
      <input
        id={id}
        name={id}
        className="p-2 border"
        defaultValue={defaultValue || ""}
      />
    </div>
  );
}
