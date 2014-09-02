package bunyan

func FilterGenerator(minStatusCode, maxStatusCode int, excludeStatusCodes []int, allowRequestBody, allowResponseBody bool) StreamFilter {
	statusCodeMap := map[int]bool{}

	for _, statusCode := range excludeStatusCodes {
		statusCodeMap[statusCode] = true
	}

	return func(entry *LogEntry) bool {
		if entry.Response != nil {
			statusCode := entry.Response.StatusCode

			if statusCode < minStatusCode || statusCode > maxStatusCode {
				return false
			}

			if statusCodeMap[statusCode] {
				return false
			}
		}

		if !allowRequestBody && entry.Request != nil {
			entry.Request.Body = nil
		}

		if !allowResponseBody && entry.Response != nil {
			entry.Response.Body = nil
		}

		return true
	}
}
