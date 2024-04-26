export type BodyType = "none" | "form-data" | "x-www-formdata" | "raw" | "binary"

export type BodyTypeItem = {
  title: string;
  value: BodyType;
}

export type KeyValueData = {
  key: string;
  value: string;
  is_active: boolean;
}

export type RowType = "text" | "file"
export type FormData = {
  key: string;
  text_value: string;
  file_value: string[];
  is_active: boolean;
  row_type: RowType;
}

export type Language = "javascript" | "xml" | "json" | "plain";
export type LanguageItem = {
  title: string;
  value: Language;
}

export type RequestTab = "options" | "header" | "body";