package webcrawler

func DropDuplicatedUrls(urls *map[string]string) {
	// TODO: 링크마다 쿼리를 날리는건 불합리함. 최근 200개 정도 링크를 메모리 상에 갖고 있고 거기 안에서 찾을까?

	for url, _ := range *urls {
		var exists bool

		err := db.QueryRow(
			"SELECT COUNT(*) FROM \"documents\" WHERE url=$1",
			url).Scan(&exists)
		if err != nil {
			panic(err)
		}

		if exists {
			delete(*urls, url)
		}
	}
}
