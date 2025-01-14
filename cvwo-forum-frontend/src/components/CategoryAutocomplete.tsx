"use client";
import { GetAllCategory } from "@/lib/getCategory";
import { Autocomplete, createFilterOptions, TextField } from "@mui/material";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { CategoryChip } from "@/lib/types/CategoryChip";

const filter = createFilterOptions<CategoryChip>();

export default function CategoryAutocomplete({
  selectedCategory,
  setSelectedCategory,
}: {
  selectedCategory: CategoryChip[];
  setSelectedCategory: Dispatch<SetStateAction<CategoryChip[]>>;
}) {
  const [chipDataList, setChipDataList] = useState<CategoryChip[]>([]);
  useEffect(() => {
    GetAllCategory().then(setChipDataList);
  }, []);
  return (
    <Autocomplete
      multiple
      handleHomeEndKeys
      value={selectedCategory}
      renderInput={(params) => (
        <TextField
          {...params}
          label="Category"
          placeholder="Type to Search or Add"
        />
      )}
      onChange={(_, val) => {
        const latestVal = val.pop();
        if (latestVal == undefined) {
          setSelectedCategory([]);
          return;
        }
        if (latestVal.id < 0) {
          setSelectedCategory([
            ...val,
            { ...latestVal, title: latestVal.title.replace("Add ", "") },
          ]);
        } else {
          setSelectedCategory([...val, latestVal]);
        }
      }}
      options={chipDataList}
      getOptionLabel={(categoryChip) => categoryChip.title}
      filterOptions={(options, params) => {
        const filtered = filter(options, params);
        const inputValue = params.inputValue;
        const isExists =
          options.some((val) => inputValue === val.title) ||
          selectedCategory.some((val) => inputValue === val.title);
        if (inputValue != "" && !isExists) {
          filtered.push({
            id: -1,
            title: `Add ${inputValue}`,
          });
        }
        return filtered;
      }}
      renderOption={(props, val) => {
        const { key, ...optionProps } = props;
        return (
          <li key={key} {...optionProps}>
            {val.title}
          </li>
        );
      }}
    />
  );
}
