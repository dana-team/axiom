import { Checkbox } from "@/components/ui/checkbox";

const NavFiltersDrawer = () => {
  const filters = ["Production", "Staging", "Development"];

  return (
    <div className="flex justify-evenly">
      {filters.map((filter, index) => (
        <label className="flex items-center gap-2">
          <Checkbox id={`filter${index}`} />
          <span>{filter}</span>
        </label>
      ))}
    </div>
  );
};

export default NavFiltersDrawer;
