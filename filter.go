package main

func FilterFormats(formats map[string][]string) {
	auxMap := make(map[string]bool, len(FormatsToDownload))

	for _, v := range FormatsToDownload {
		auxMap[v] = true
	}

	for k := range formats {
		if _, ok := auxMap[k]; !ok {
			delete(formats, k)
		}
	}
}
