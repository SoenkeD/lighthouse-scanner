package lib

import "time"

type LightHouseResp struct {
	LighthouseVersion string    `json:"lighthouseVersion"`
	RequestedUrl      string    `json:"requestedUrl"`
	FinalUrl          string    `json:"finalUrl"`
	FetchTime         time.Time `json:"fetchTime"`
	GatherMode        string    `json:"gatherMode"`
	RunWarnings       []string  `json:"runWarnings"`
	UserAgent         string    `json:"userAgent"`
	Environment       struct {
		NetworkUserAgent string  `json:"networkUserAgent"`
		HostUserAgent    string  `json:"hostUserAgent"`
		BenchmarkIndex   float64 `json:"benchmarkIndex"`
		Credits          struct {
			AxeCore string `json:"axe-core"`
		} `json:"credits"`
	} `json:"environment"`
	Audits struct {
		IsOnHttps struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"is-on-https"`
		ServiceWorker struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
		} `json:"service-worker"`
		Viewport struct {
			Id               string        `json:"id"`
			Title            string        `json:"title"`
			Description      string        `json:"description"`
			Score            float64       `json:"score"`
			ScoreDisplayMode string        `json:"scoreDisplayMode"`
			Warnings         []interface{} `json:"warnings"`
		} `json:"viewport"`
		FirstContentfulPaint struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
		} `json:"first-contentful-paint"`
		LargestContentfulPaint struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
		} `json:"largest-contentful-paint"`
		FirstMeaningfulPaint struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
		} `json:"first-meaningful-paint"`
		SpeedIndex struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
		} `json:"speed-index"`
		ScreenshotThumbnails struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type  string  `json:"type"`
				Scale float64 `json:"scale"`
				Items []struct {
					Timing    float64 `json:"timing"`
					Timestamp float64 `json:"timestamp"`
					Data      string  `json:"data"`
				} `json:"items"`
			} `json:"details"`
		} `json:"screenshot-thumbnails"`
		FinalScreenshot struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type      string  `json:"type"`
				Timing    float64 `json:"timing"`
				Timestamp float64 `json:"timestamp"`
				Data      string  `json:"data"`
			} `json:"details"`
		} `json:"final-screenshot"`
		TotalBlockingTime struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
		} `json:"total-blocking-time"`
		MaxPotentialFid struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
		} `json:"max-potential-fid"`
		CumulativeLayoutShift struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type  string `json:"type"`
				Items []struct {
					CumulativeLayoutShiftMainFrame float64 `json:"cumulativeLayoutShiftMainFrame"`
					TotalCumulativeLayoutShift     float64 `json:"totalCumulativeLayoutShift"`
				} `json:"items"`
			} `json:"details"`
		} `json:"cumulative-layout-shift"`
		ErrorsInConsole struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"errors-in-console"`
		ServerResponseTime struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key       string `json:"key"`
					ValueType string `json:"valueType"`
					Label     string `json:"label"`
				} `json:"headings"`
				Items []struct {
					Url          string  `json:"url"`
					ResponseTime float64 `json:"responseTime"`
				} `json:"items"`
				OverallSavingsMs float64 `json:"overallSavingsMs"`
			} `json:"details"`
		} `json:"server-response-time"`
		Interactive struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
		} `json:"interactive"`
		UserTimings struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"user-timings"`
		CriticalRequestChains struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			DisplayValue     string      `json:"displayValue"`
			Details          struct {
				Type   string `json:"type"`
				Chains struct {
					EF6732A08D0C2075A0417DB6946AF589 struct {
						Request struct {
							Url                  string  `json:"url"`
							StartTime            float64 `json:"startTime"`
							EndTime              float64 `json:"endTime"`
							ResponseReceivedTime float64 `json:"responseReceivedTime"`
							TransferSize         float64 `json:"transferSize"`
						} `json:"request"`
						Children struct {
							Field1 struct {
								Request struct {
									Url                  string  `json:"url"`
									StartTime            float64 `json:"startTime"`
									EndTime              float64 `json:"endTime"`
									ResponseReceivedTime float64 `json:"responseReceivedTime"`
									TransferSize         float64 `json:"transferSize"`
								} `json:"request"`
							} `json:"27197.2"`
							Field2 struct {
								Request struct {
									Url                  string  `json:"url"`
									StartTime            float64 `json:"startTime"`
									EndTime              float64 `json:"endTime"`
									ResponseReceivedTime float64 `json:"responseReceivedTime"`
									TransferSize         float64 `json:"transferSize"`
								} `json:"request"`
							} `json:"27197.3"`
							Field3 struct {
								Request struct {
									Url                  string  `json:"url"`
									StartTime            float64 `json:"startTime"`
									EndTime              float64 `json:"endTime"`
									ResponseReceivedTime float64 `json:"responseReceivedTime"`
									TransferSize         float64 `json:"transferSize"`
								} `json:"request"`
							} `json:"27197.4"`
							Field4 struct {
								Request struct {
									Url                  string  `json:"url"`
									StartTime            float64 `json:"startTime"`
									EndTime              float64 `json:"endTime"`
									ResponseReceivedTime float64 `json:"responseReceivedTime"`
									TransferSize         float64 `json:"transferSize"`
								} `json:"request"`
							} `json:"27197.5"`
							Field5 struct {
								Request struct {
									Url                  string  `json:"url"`
									StartTime            float64 `json:"startTime"`
									EndTime              float64 `json:"endTime"`
									ResponseReceivedTime float64 `json:"responseReceivedTime"`
									TransferSize         float64 `json:"transferSize"`
								} `json:"request"`
							} `json:"27197.6"`
							Field6 struct {
								Request struct {
									Url                  string  `json:"url"`
									StartTime            float64 `json:"startTime"`
									EndTime              float64 `json:"endTime"`
									ResponseReceivedTime float64 `json:"responseReceivedTime"`
									TransferSize         float64 `json:"transferSize"`
								} `json:"request"`
							} `json:"27197.7"`
							Field7 struct {
								Request struct {
									Url                  string  `json:"url"`
									StartTime            float64 `json:"startTime"`
									EndTime              float64 `json:"endTime"`
									ResponseReceivedTime float64 `json:"responseReceivedTime"`
									TransferSize         float64 `json:"transferSize"`
								} `json:"request"`
							} `json:"27197.8"`
							Field8 struct {
								Request struct {
									Url                  string  `json:"url"`
									StartTime            float64 `json:"startTime"`
									EndTime              float64 `json:"endTime"`
									ResponseReceivedTime float64 `json:"responseReceivedTime"`
									TransferSize         float64 `json:"transferSize"`
								} `json:"request"`
							} `json:"27197.9"`
						} `json:"children"`
					} `json:"EF6732A08D0C2075A0417DB6946AF589"`
				} `json:"chains"`
				LongestChain struct {
					Duration     float64 `json:"duration"`
					Length       float64 `json:"length"`
					TransferSize float64 `json:"transferSize"`
				} `json:"longestChain"`
			} `json:"details"`
		} `json:"critical-request-chains"`
		Redirects struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type             string        `json:"type"`
				Headings         []interface{} `json:"headings"`
				Items            []interface{} `json:"items"`
				OverallSavingsMs float64       `json:"overallSavingsMs"`
			} `json:"details"`
		} `json:"redirects"`
		InstallableManifest struct {
			Id               string        `json:"id"`
			Title            string        `json:"title"`
			Description      string        `json:"description"`
			Score            float64       `json:"score"`
			ScoreDisplayMode string        `json:"scoreDisplayMode"`
			Warnings         []interface{} `json:"warnings"`
			Details          struct {
				Type      string        `json:"type"`
				Headings  []interface{} `json:"headings"`
				Items     []interface{} `json:"items"`
				DebugData struct {
					Type        string `json:"type"`
					ManifestUrl string `json:"manifestUrl"`
				} `json:"debugData"`
			} `json:"details"`
		} `json:"installable-manifest"`
		AppleTouchIcon struct {
			Id               string        `json:"id"`
			Title            string        `json:"title"`
			Description      string        `json:"description"`
			Score            float64       `json:"score"`
			ScoreDisplayMode string        `json:"scoreDisplayMode"`
			Warnings         []interface{} `json:"warnings"`
		} `json:"apple-touch-icon"`
		SplashScreen struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Explanation      string  `json:"explanation"`
			Details          struct {
				Type  string `json:"type"`
				Items []struct {
					Failures             []string `json:"failures"`
					IsParseFailure       bool     `json:"isParseFailure"`
					HasStartUrl          bool     `json:"hasStartUrl"`
					HasIconsAtLeast144Px bool     `json:"hasIconsAtLeast144px"`
					HasIconsAtLeast512Px bool     `json:"hasIconsAtLeast512px"`
					FetchesIcon          bool     `json:"fetchesIcon"`
					HasPWADisplayValue   bool     `json:"hasPWADisplayValue"`
					HasBackgroundColor   bool     `json:"hasBackgroundColor"`
					HasThemeColor        bool     `json:"hasThemeColor"`
					HasShortName         bool     `json:"hasShortName"`
					ShortNameLength      bool     `json:"shortNameLength"`
					HasName              bool     `json:"hasName"`
					HasMaskableIcon      bool     `json:"hasMaskableIcon"`
				} `json:"items"`
			} `json:"details"`
		} `json:"splash-screen"`
		ThemedOmnibox struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Explanation      string  `json:"explanation"`
			Details          struct {
				Type  string `json:"type"`
				Items []struct {
					Failures             []string `json:"failures"`
					ThemeColor           string   `json:"themeColor"`
					IsParseFailure       bool     `json:"isParseFailure"`
					HasStartUrl          bool     `json:"hasStartUrl"`
					HasIconsAtLeast144Px bool     `json:"hasIconsAtLeast144px"`
					HasIconsAtLeast512Px bool     `json:"hasIconsAtLeast512px"`
					FetchesIcon          bool     `json:"fetchesIcon"`
					HasPWADisplayValue   bool     `json:"hasPWADisplayValue"`
					HasBackgroundColor   bool     `json:"hasBackgroundColor"`
					HasThemeColor        bool     `json:"hasThemeColor"`
					HasShortName         bool     `json:"hasShortName"`
					ShortNameLength      bool     `json:"shortNameLength"`
					HasName              bool     `json:"hasName"`
					HasMaskableIcon      bool     `json:"hasMaskableIcon"`
				} `json:"items"`
			} `json:"details"`
		} `json:"themed-omnibox"`
		MaskableIcon struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
		} `json:"maskable-icon"`
		ContentWidth struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
		} `json:"content-width"`
		ImageAspectRatio struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"image-aspect-ratio"`
		ImageSizeResponsive struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"image-size-responsive"`
		PreloadFonts struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"preload-fonts"`
		Deprecations struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key      string `json:"key"`
					ItemType string `json:"itemType"`
					Text     string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Value  string `json:"value"`
					Source struct {
						Type        string  `json:"type"`
						Url         string  `json:"url"`
						UrlProvider string  `json:"urlProvider"`
						Line        float64 `json:"line"`
						Column      float64 `json:"column"`
					} `json:"source"`
				} `json:"items"`
			} `json:"details"`
		} `json:"deprecations"`
		MainthreadWorkBreakdown struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key         string  `json:"key"`
					ItemType    string  `json:"itemType"`
					Text        string  `json:"text"`
					Granularity float64 `json:"granularity,omitempty"`
				} `json:"headings"`
				Items []struct {
					Group      string  `json:"group"`
					GroupLabel string  `json:"groupLabel"`
					Duration   float64 `json:"duration"`
				} `json:"items"`
			} `json:"details"`
		} `json:"mainthread-work-breakdown"`
		BootupTime struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key         string  `json:"key"`
					ItemType    string  `json:"itemType"`
					Text        string  `json:"text"`
					Granularity float64 `json:"granularity,omitempty"`
				} `json:"headings"`
				Items []struct {
					Url                string  `json:"url"`
					Total              float64 `json:"total"`
					Scripting          float64 `json:"scripting"`
					ScriptParseCompile float64 `json:"scriptParseCompile"`
				} `json:"items"`
				Summary struct {
					WastedMs float64 `json:"wastedMs"`
				} `json:"summary"`
			} `json:"details"`
		} `json:"bootup-time"`
		UsesRelPreload struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type             string        `json:"type"`
				Headings         []interface{} `json:"headings"`
				Items            []interface{} `json:"items"`
				OverallSavingsMs float64       `json:"overallSavingsMs"`
			} `json:"details"`
		} `json:"uses-rel-preload"`
		UsesRelPreconnect struct {
			Id               string        `json:"id"`
			Title            string        `json:"title"`
			Description      string        `json:"description"`
			score            float64       `json:"score"`
			ScoreDisplayMode string        `json:"scoreDisplayMode"`
			Warnings         []interface{} `json:"warnings"`
		} `json:"uses-rel-preconnect"`
		FontDisplay struct {
			Id               string        `json:"id"`
			Title            string        `json:"title"`
			Description      string        `json:"description"`
			score            float64       `json:"score"`
			ScoreDisplayMode string        `json:"scoreDisplayMode"`
			Warnings         []interface{} `json:"warnings"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"font-display"`
		Diagnostics struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type  string `json:"type"`
				Items []struct {
					NumRequests              float64 `json:"numRequests"`
					NumScripts               float64 `json:"numScripts"`
					NumStylesheets           float64 `json:"numStylesheets"`
					NumFonts                 float64 `json:"numFonts"`
					NumTasks                 float64 `json:"numTasks"`
					NumTasksOver10Ms         float64 `json:"numTasksOver10ms"`
					NumTasksOver25Ms         float64 `json:"numTasksOver25ms"`
					NumTasksOver50Ms         float64 `json:"numTasksOver50ms"`
					NumTasksOver100Ms        float64 `json:"numTasksOver100ms"`
					NumTasksOver500Ms        float64 `json:"numTasksOver500ms"`
					Rtt                      float64 `json:"rtt"`
					Throughput               float64 `json:"throughput"`
					MaxRtt                   float64 `json:"maxRtt"`
					MaxServerLatency         float64 `json:"maxServerLatency"`
					TotalByteWeight          float64 `json:"totalByteWeight"`
					TotalTaskTime            float64 `json:"totalTaskTime"`
					MainDocumentTransferSize float64 `json:"mainDocumentTransferSize"`
				} `json:"items"`
			} `json:"details"`
		} `json:"diagnostics"`
		NetworkRequests struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key         string  `json:"key"`
					ItemType    string  `json:"itemType"`
					Text        string  `json:"text"`
					Granularity float64 `json:"granularity,omitempty"`
					DisplayUnit string  `json:"displayUnit,omitempty"`
				} `json:"headings"`
				Items []struct {
					Url                       string  `json:"url"`
					Protocol                  string  `json:"protocol"`
					StartTime                 float64 `json:"startTime"`
					EndTime                   float64 `json:"endTime"`
					Finished                  bool    `json:"finished"`
					TransferSize              float64 `json:"transferSize"`
					ResourceSize              float64 `json:"resourceSize"`
					StatusCode                float64 `json:"statusCode"`
					MimeType                  string  `json:"mimeType"`
					ResourceType              string  `json:"resourceType"`
					Priority                  string  `json:"priority"`
					ExperimentalFromMainFrame bool    `json:"experimentalFromMainFrame"`
					IsLinkPreload             bool    `json:"isLinkPreload,omitempty"`
				} `json:"items"`
				DebugData struct {
					Type               string  `json:"type"`
					NetworkStartTimeTs float64 `json:"networkStartTimeTs"`
				} `json:"debugData"`
			} `json:"details"`
		} `json:"network-requests"`
		NetworkRtt struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			NumericValue     float64     `json:"numericValue"`
			NumericUnit      string      `json:"numericUnit"`
			DisplayValue     string      `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key         string  `json:"key"`
					ItemType    string  `json:"itemType"`
					Text        string  `json:"text"`
					Granularity float64 `json:"granularity,omitempty"`
				} `json:"headings"`
				Items []struct {
					Origin string  `json:"origin"`
					Rtt    float64 `json:"rtt"`
				} `json:"items"`
			} `json:"details"`
		} `json:"network-rtt"`
		NetworkServerLatency struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			NumericValue     float64     `json:"numericValue"`
			NumericUnit      string      `json:"numericUnit"`
			DisplayValue     string      `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key         string  `json:"key"`
					ItemType    string  `json:"itemType"`
					Text        string  `json:"text"`
					Granularity float64 `json:"granularity,omitempty"`
				} `json:"headings"`
				Items []struct {
					Origin             string  `json:"origin"`
					ServerResponseTime float64 `json:"serverResponseTime"`
				} `json:"items"`
			} `json:"details"`
		} `json:"network-server-latency"`
		MainThreadTasks struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key         string  `json:"key"`
					ItemType    string  `json:"itemType"`
					Granularity float64 `json:"granularity"`
					Text        string  `json:"text"`
				} `json:"headings"`
				Items []struct {
					Duration  float64 `json:"duration"`
					StartTime float64 `json:"startTime"`
				} `json:"items"`
			} `json:"details"`
		} `json:"main-thread-tasks"`
		Metrics struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			numericValue     float64     `json:"numericValue"`
			NumericUnit      string      `json:"numericUnit"`
			Details          struct {
				Type  string `json:"type"`
				Items []struct {
					FirstContentfulPaint                      float64 `json:"firstContentfulPaint,omitempty"`
					FirstMeaningfulPaint                      float64 `json:"firstMeaningfulPaint,omitempty"`
					LargestContentfulPaint                    float64 `json:"largestContentfulPaint,omitempty"`
					Interactive                               float64 `json:"interactive,omitempty"`
					SpeedIndex                                float64 `json:"speedIndex,omitempty"`
					TotalBlockingTime                         float64 `json:"totalBlockingTime,omitempty"`
					MaxPotentialFID                           float64 `json:"maxPotentialFID,omitempty"`
					CumulativeLayoutShift                     float64 `json:"cumulativeLayoutShift,omitempty"`
					CumulativeLayoutShiftMainFrame            float64 `json:"cumulativeLayoutShiftMainFrame,omitempty"`
					TotalCumulativeLayoutShift                float64 `json:"totalCumulativeLayoutShift,omitempty"`
					ObservedTimeOrigin                        float64 `json:"observedTimeOrigin,omitempty"`
					ObservedTimeOriginTs                      int64   `json:"observedTimeOriginTs,omitempty"`
					ObservedNavigationStart                   float64 `json:"observedNavigationStart,omitempty"`
					ObservedNavigationStartTs                 int64   `json:"observedNavigationStartTs,omitempty"`
					ObservedFirstPaint                        float64 `json:"observedFirstPaint,omitempty"`
					ObservedFirstPaintTs                      int64   `json:"observedFirstPaintTs,omitempty"`
					ObservedFirstContentfulPaint              float64 `json:"observedFirstContentfulPaint,omitempty"`
					ObservedFirstContentfulPaintTs            int64   `json:"observedFirstContentfulPaintTs,omitempty"`
					ObservedFirstContentfulPaintAllFrames     float64 `json:"observedFirstContentfulPaintAllFrames,omitempty"`
					ObservedFirstContentfulPaintAllFramesTs   int64   `json:"observedFirstContentfulPaintAllFramesTs,omitempty"`
					ObservedFirstMeaningfulPaint              float64 `json:"observedFirstMeaningfulPaint,omitempty"`
					ObservedFirstMeaningfulPaintTs            int64   `json:"observedFirstMeaningfulPaintTs,omitempty"`
					ObservedLargestContentfulPaint            float64 `json:"observedLargestContentfulPaint,omitempty"`
					ObservedLargestContentfulPaintTs          int64   `json:"observedLargestContentfulPaintTs,omitempty"`
					ObservedLargestContentfulPaintAllFrames   float64 `json:"observedLargestContentfulPaintAllFrames,omitempty"`
					ObservedLargestContentfulPaintAllFramesTs int64   `json:"observedLargestContentfulPaintAllFramesTs,omitempty"`
					ObservedTraceEnd                          float64 `json:"observedTraceEnd,omitempty"`
					ObservedTraceEndTs                        int64   `json:"observedTraceEndTs,omitempty"`
					ObservedLoad                              float64 `json:"observedLoad,omitempty"`
					ObservedLoadTs                            int64   `json:"observedLoadTs,omitempty"`
					ObservedDomContentLoaded                  float64 `json:"observedDomContentLoaded,omitempty"`
					ObservedDomContentLoadedTs                int64   `json:"observedDomContentLoadedTs,omitempty"`
					ObservedCumulativeLayoutShift             float64 `json:"observedCumulativeLayoutShift,omitempty"`
					ObservedCumulativeLayoutShiftMainFrame    float64 `json:"observedCumulativeLayoutShiftMainFrame,omitempty"`
					ObservedTotalCumulativeLayoutShift        float64 `json:"observedTotalCumulativeLayoutShift,omitempty"`
					ObservedFirstVisualChange                 float64 `json:"observedFirstVisualChange,omitempty"`
					ObservedFirstVisualChangeTs               int64   `json:"observedFirstVisualChangeTs,omitempty"`
					ObservedLastVisualChange                  float64 `json:"observedLastVisualChange,omitempty"`
					ObservedLastVisualChangeTs                int64   `json:"observedLastVisualChangeTs,omitempty"`
					ObservedSpeedIndex                        float64 `json:"observedSpeedIndex,omitempty"`
					ObservedSpeedIndexTs                      int64   `json:"observedSpeedIndexTs,omitempty"`
					LcpInvalidated                            bool    `json:"lcpInvalidated,omitempty"`
				} `json:"items"`
			} `json:"details"`
		} `json:"metrics"`
		PerformanceBudget struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"performance-budget"`
		TimingBudget struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"timing-budget"`
		ResourceSummary struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			DisplayValue     string      `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key      string `json:"key"`
					ItemType string `json:"itemType"`
					Text     string `json:"text"`
				} `json:"headings"`
				Items []struct {
					ResourceType string  `json:"resourceType"`
					Label        string  `json:"label"`
					RequestCount float64 `json:"requestCount"`
					TransferSize float64 `json:"transferSize"`
				} `json:"items"`
			} `json:"details"`
		} `json:"resource-summary"`
		ThirdPartySummary struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"third-party-summary"`
		ThirdPartyFacades struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"third-party-facades"`
		LargestContentfulPaintElement struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			DisplayValue     string      `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key      string `json:"key"`
					ItemType string `json:"itemType"`
					Text     string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet   string `json:"snippet"`
						NodeLabel string `json:"nodeLabel"`
					} `json:"node"`
				} `json:"items"`
			} `json:"details"`
		} `json:"largest-contentful-paint-element"`
		LcpLazyLoaded struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key      string `json:"key"`
					ItemType string `json:"itemType"`
					Text     string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet   string `json:"snippet"`
						NodeLabel string `json:"nodeLabel"`
					} `json:"node"`
				} `json:"items"`
			} `json:"details"`
		} `json:"lcp-lazy-loaded"`
		LayoutShiftElements struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			DisplayValue     string      `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key         string  `json:"key"`
					ItemType    string  `json:"itemType"`
					Text        string  `json:"text"`
					Granularity float64 `json:"granularity,omitempty"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet   string `json:"snippet"`
						NodeLabel string `json:"nodeLabel"`
					} `json:"node"`
					Score float64 `json:"score"`
				} `json:"items"`
			} `json:"details"`
		} `json:"layout-shift-elements"`
		LongTasks struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			DisplayValue     string      `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key         string  `json:"key"`
					ItemType    string  `json:"itemType"`
					Text        string  `json:"text"`
					Granularity float64 `json:"granularity,omitempty"`
				} `json:"headings"`
				Items []struct {
					Url       string  `json:"url"`
					Duration  float64 `json:"duration"`
					StartTime float64 `json:"startTime"`
				} `json:"items"`
			} `json:"details"`
		} `json:"long-tasks"`
		NoUnloadListeners struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
		} `json:"no-unload-listeners"`
		NonCompositedAnimations struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"non-composited-animations"`
		UnsizedImages struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"unsized-images"`
		ValidSourceMaps struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key             string `json:"key"`
					ItemType        string `json:"itemType"`
					SubItemsHeading struct {
						Key string `json:"key"`
					} `json:"subItemsHeading,omitempty"`
					Text string `json:"text"`
				} `json:"headings"`
				Items []struct {
					ScriptUrl    string `json:"scriptUrl"`
					SourceMapUrl string `json:"sourceMapUrl"`
					SubItems     struct {
						Type  string `json:"type"`
						Items []struct {
							Error string `json:"error"`
						} `json:"items"`
					} `json:"subItems"`
				} `json:"items"`
			} `json:"details"`
		} `json:"valid-source-maps"`
		PreloadLcpImage struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key       string `json:"key"`
					ValueType string `json:"valueType"`
					Label     string `json:"label"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet   string `json:"snippet"`
						NodeLabel string `json:"nodeLabel"`
					} `json:"node"`
					Url      string  `json:"url"`
					WastedMs float64 `json:"wastedMs"`
				} `json:"items"`
				OverallSavingsMs float64 `json:"overallSavingsMs"`
				DebugData        struct {
					Type          string `json:"type"`
					InitiatorPath []struct {
						Url           string `json:"url"`
						InitiatorType string `json:"initiatorType"`
					} `json:"initiatorPath"`
					PathLength float64 `json:"pathLength"`
				} `json:"debugData"`
			} `json:"details"`
		} `json:"preload-lcp-image"`
		CspXss struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key             string `json:"key"`
					ItemType        string `json:"itemType"`
					SubItemsHeading struct {
						Key string `json:"key"`
					} `json:"subItemsHeading"`
					Text string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Directive   string `json:"directive"`
					Description string `json:"description"`
					Severity    string `json:"severity"`
				} `json:"items"`
			} `json:"details"`
		} `json:"csp-xss"`
		FullPageScreenshot struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type       string `json:"type"`
				Screenshot struct {
					Data   string  `json:"data"`
					Width  float64 `json:"width"`
					Height float64 `json:"height"`
				} `json:"screenshot"`
				Nodes struct {
					Page0IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-0-IMG"`
					Page1IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-1-IMG"`
					Page2IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-2-IMG"`
					Page3IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-3-IMG"`
					Page4IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-4-IMG"`
					Page5IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-5-IMG"`
					Page6IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-6-IMG"`
					Page7IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-7-IMG"`
					Page8IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-8-IMG"`
					Page9IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-9-IMG"`
					Page10IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-10-IMG"`
					Page11IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-11-IMG"`
					Page12IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-12-IMG"`
					Page13IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-13-IMG"`
					Page14IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-14-IMG"`
					Page15IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-15-IMG"`
					Page16IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-16-IMG"`
					Page17IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-17-IMG"`
					Page18IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-18-IMG"`
					Page19IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-19-IMG"`
					Page20IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-20-IMG"`
					Page21IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-21-IMG"`
					Page22IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-22-IMG"`
					Page23IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-23-IMG"`
					Page24IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-24-IMG"`
					Page25IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-25-IMG"`
					Page26IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-26-IMG"`
					Page27IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-27-IMG"`
					Page28IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-28-IMG"`
					Page29IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-29-IMG"`
					Page30IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-30-IMG"`
					Page31IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-31-IMG"`
					Page32IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-32-IMG"`
					Page33IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-33-IMG"`
					Page34IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-34-IMG"`
					Page35IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-35-IMG"`
					Page36IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-36-IMG"`
					Page37IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-37-IMG"`
					Page38IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-38-IMG"`
					Page39IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-39-IMG"`
					Page40IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-40-IMG"`
					Page41IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-41-IMG"`
					Page42IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-42-IMG"`
					Page43IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-43-IMG"`
					Page44IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-44-IMG"`
					Page45IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-45-IMG"`
					Page46IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-46-IMG"`
					Page47IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-47-IMG"`
					Page48IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-48-IMG"`
					Page49IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-49-IMG"`
					Page50IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-50-IMG"`
					Page51IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-51-IMG"`
					Page52IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-52-IMG"`
					Page53IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-53-IMG"`
					Page54IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-54-IMG"`
					Page55IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-55-IMG"`
					Page56IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-56-IMG"`
					Page57IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-57-IMG"`
					Page58IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-58-IMG"`
					Page59IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-59-IMG"`
					Page60IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-60-IMG"`
					Page61IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-61-IMG"`
					Page62IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-62-IMG"`
					Page63IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-63-IMG"`
					Page64IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-64-IMG"`
					Page65IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-65-IMG"`
					Page66IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-66-IMG"`
					Page67IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-67-IMG"`
					Page68IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-68-IMG"`
					Page69IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-69-IMG"`
					Page70IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-70-IMG"`
					Page71IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-71-IMG"`
					Page72IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-72-IMG"`
					Page73IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-73-IMG"`
					Page74IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-74-IMG"`
					Page75IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-75-IMG"`
					Page76IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-76-IMG"`
					Page77IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-77-IMG"`
					Page78IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-78-IMG"`
					Page79IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-79-IMG"`
					Page80IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-80-IMG"`
					Page81IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-81-IMG"`
					Page82IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-82-IMG"`
					Page83IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-83-IMG"`
					Page84IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-84-IMG"`
					Page85IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-85-IMG"`
					Page86IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-86-IMG"`
					Page87IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-87-IMG"`
					Page88IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-88-IMG"`
					Page89IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-89-IMG"`
					Page90IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-90-IMG"`
					Page91IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-91-IMG"`
					Page92IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-92-IMG"`
					Page93IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-93-IMG"`
					Page94IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-94-IMG"`
					Page95IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-95-IMG"`
					Page96IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-96-IMG"`
					Page97IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-97-IMG"`
					Page98IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-98-IMG"`
					Page99IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-99-IMG"`
					Page100IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-100-IMG"`
					Page101IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-101-IMG"`
					Page102IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-102-IMG"`
					Page103IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-103-IMG"`
					Page104IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-104-IMG"`
					Page105IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-105-IMG"`
					Page106IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-106-IMG"`
					Page107IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-107-IMG"`
					Page108IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-108-IMG"`
					Page109IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-109-IMG"`
					Page110IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-110-IMG"`
					Page111IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-111-IMG"`
					Page112IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-112-IMG"`
					Page113IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-113-IMG"`
					Page114IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-114-IMG"`
					Page115IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-115-IMG"`
					Page116IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-116-IMG"`
					Page117IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-117-IMG"`
					Page118IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-118-IMG"`
					Page119IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-119-IMG"`
					Page120IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-120-IMG"`
					Page121IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-121-IMG"`
					Page122IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-122-IMG"`
					Page123IMG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-123-IMG"`
					Page124FORM struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-124-FORM"`
					Page125P struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-125-P"`
					Page126DIV struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-126-DIV"`
					Page127H1 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-127-H1"`
					Page128SPAN struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-128-SPAN"`
					Page129SPAN struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-129-SPAN"`
					Page130SPAN struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-130-SPAN"`
					Page131SPAN struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-131-SPAN"`
					Page132SPAN struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-132-SPAN"`
					Page133SPAN struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-133-SPAN"`
					Page134SPAN struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-134-SPAN"`
					Page135SPAN struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-135-SPAN"`
					Page136CANVAS struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-136-CANVAS"`
					Page137Svg struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"page-137-svg"`
					A struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-0-A"`
					A1 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-1-A"`
					A2 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-2-A"`
					A3 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-3-A"`
					A4 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-4-A"`
					A5 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-5-A"`
					A6 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-6-A"`
					A7 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-7-A"`
					A8 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-8-A"`
					A9 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-9-A"`
					A10 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-10-A"`
					A11 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-11-A"`
					A12 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-12-A"`
					A13 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-13-A"`
					A14 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-14-A"`
					A15 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-15-A"`
					A16 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-16-A"`
					A17 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-17-A"`
					A18 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-18-A"`
					A19 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-19-A"`
					A20 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-20-A"`
					A21 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-21-A"`
					A22 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-22-A"`
					A23 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-23-A"`
					A24 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-24-A"`
					A25 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-25-A"`
					A26 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-26-A"`
					A27 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-27-A"`
					A28 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-28-A"`
					A29 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-29-A"`
					A30 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-30-A"`
					A31 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-31-A"`
					A32 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-32-A"`
					A33 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-33-A"`
					A34 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-34-A"`
					A35 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-35-A"`
					A36 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-36-A"`
					A37 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-37-A"`
					A38 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-38-A"`
					A39 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-39-A"`
					A40 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-40-A"`
					A41 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-41-A"`
					A42 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-42-A"`
					A43 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-43-A"`
					A44 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-44-A"`
					A45 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-45-A"`
					A46 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-46-A"`
					A47 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-47-A"`
					A48 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-48-A"`
					A49 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-49-A"`
					A50 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-50-A"`
					A51 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-51-A"`
					A52 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-52-A"`
					A53 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-53-A"`
					A54 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-54-A"`
					A55 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-55-A"`
					A56 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-56-A"`
					A57 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-57-A"`
					A58 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-58-A"`
					A59 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-59-A"`
					A60 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-60-A"`
					A61 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-61-A"`
					A62 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-62-A"`
					A63 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-63-A"`
					A64 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-64-A"`
					A65 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-65-A"`
					A66 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-66-A"`
					A67 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-67-A"`
					A68 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-68-A"`
					A69 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-69-A"`
					A70 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-70-A"`
					A71 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-71-A"`
					A72 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-72-A"`
					A73 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-73-A"`
					A74 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-74-A"`
					A75 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-75-A"`
					A76 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-76-A"`
					A77 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-77-A"`
					A78 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-78-A"`
					A79 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-79-A"`
					A80 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-80-A"`
					A81 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-81-A"`
					A82 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-82-A"`
					A83 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-83-A"`
					A84 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-84-A"`
					A85 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-85-A"`
					A86 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-86-A"`
					A87 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-87-A"`
					A88 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-88-A"`
					A89 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-89-A"`
					A90 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-90-A"`
					A91 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-91-A"`
					A92 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-92-A"`
					A93 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-93-A"`
					A94 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-94-A"`
					A95 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-95-A"`
					A96 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-96-A"`
					A97 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-97-A"`
					A98 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-98-A"`
					A99 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-99-A"`
					A100 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-100-A"`
					A101 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-101-A"`
					A102 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-102-A"`
					A103 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-103-A"`
					A104 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-104-A"`
					A105 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-105-A"`
					A106 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-106-A"`
					A107 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-107-A"`
					A108 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-108-A"`
					A109 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-109-A"`
					A110 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-110-A"`
					A111 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-111-A"`
					A112 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-112-A"`
					A113 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-113-A"`
					A114 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-114-A"`
					A115 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-115-A"`
					LINK struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-116-LINK"`
					LINK1 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-117-LINK"`
					LINK2 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-118-LINK"`
					LINK3 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-119-LINK"`
					LINK4 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-120-LINK"`
					LINK5 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-121-LINK"`
					LINK6 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-122-LINK"`
					LINK7 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-123-LINK"`
					LINK8 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-124-LINK"`
					LINK9 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-125-LINK"`
					LINK10 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-126-LINK"`
					LINK11 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-127-LINK"`
					LINK12 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-128-LINK"`
					LINK13 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-129-LINK"`
					LINK14 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-130-LINK"`
					LINK15 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-131-LINK"`
					LINK16 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-132-LINK"`
					LINK17 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-133-LINK"`
					LINK18 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-134-LINK"`
					LINK19 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-135-LINK"`
					LINK20 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-136-LINK"`
					LINK21 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-137-LINK"`
					LINK22 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-138-LINK"`
					LINK23 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-139-LINK"`
					LINK24 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-140-LINK"`
					LINK25 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-141-LINK"`
					LINK26 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-142-LINK"`
					LINK27 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-143-LINK"`
					LINK28 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-144-LINK"`
					LINK29 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-145-LINK"`
					LINK30 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-146-LINK"`
					LINK31 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-147-LINK"`
					LINK32 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-148-LINK"`
					LINK33 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-149-LINK"`
					LINK34 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-150-LINK"`
					META struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-151-META"`
					META1 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-152-META"`
					META2 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-153-META"`
					META3 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-154-META"`
					META4 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-155-META"`
					META5 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-156-META"`
					META6 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-157-META"`
					META7 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-158-META"`
					META8 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-159-META"`
					META9 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-160-META"`
					META10 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-161-META"`
					META11 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-162-META"`
					META12 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-163-META"`
					META13 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-164-META"`
					META14 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-165-META"`
					META15 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-166-META"`
					META16 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-167-META"`
					META17 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-168-META"`
					META18 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-169-META"`
					META19 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-170-META"`
					META20 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-171-META"`
					META21 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-172-META"`
					META22 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-173-META"`
					META23 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-174-META"`
					META24 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-175-META"`
					META25 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-176-META"`
					META26 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-177-META"`
					META27 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-178-META"`
					META28 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-179-META"`
					META29 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-180-META"`
					META30 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-181-META"`
					META31 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-182-META"`
					META32 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-183-META"`
					META33 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-184-META"`
					META34 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-185-META"`
					META35 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-186-META"`
					META36 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-187-META"`
					META37 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-188-META"`
					META38 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-189-META"`
					META39 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-190-META"`
					META40 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-191-META"`
					META41 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-192-META"`
					META42 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-193-META"`
					META43 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-194-META"`
					META44 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-195-META"`
					META45 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-196-META"`
					META46 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-197-META"`
					META47 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-198-META"`
					META48 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-199-META"`
					SCRIPT struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-200-SCRIPT"`
					SCRIPT1 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-201-SCRIPT"`
					SCRIPT2 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-202-SCRIPT"`
					SCRIPT3 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-203-SCRIPT"`
					SCRIPT4 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-204-SCRIPT"`
					SCRIPT5 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-205-SCRIPT"`
					SCRIPT6 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-206-SCRIPT"`
					SCRIPT7 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-207-SCRIPT"`
					SCRIPT8 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-208-SCRIPT"`
					SCRIPT9 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-209-SCRIPT"`
					SCRIPT10 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-210-SCRIPT"`
					SCRIPT11 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-211-SCRIPT"`
					SCRIPT12 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-212-SCRIPT"`
					SCRIPT13 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-213-SCRIPT"`
					SCRIPT14 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-214-SCRIPT"`
					SCRIPT15 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-215-SCRIPT"`
					SCRIPT16 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-216-SCRIPT"`
					SCRIPT17 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-217-SCRIPT"`
					SCRIPT18 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-218-SCRIPT"`
					SCRIPT19 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-219-SCRIPT"`
					SCRIPT20 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-220-SCRIPT"`
					SCRIPT21 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-221-SCRIPT"`
					SCRIPT22 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-222-SCRIPT"`
					SCRIPT23 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-223-SCRIPT"`
					SCRIPT24 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-224-SCRIPT"`
					SCRIPT25 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-225-SCRIPT"`
					SCRIPT26 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-226-SCRIPT"`
					SCRIPT27 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-227-SCRIPT"`
					SCRIPT28 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-228-SCRIPT"`
					SCRIPT29 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-229-SCRIPT"`
					SCRIPT30 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-230-SCRIPT"`
					SCRIPT31 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-231-SCRIPT"`
					SCRIPT32 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-232-SCRIPT"`
					FORM struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-233-FORM"`
					FORM1 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-234-FORM"`
					LABEL struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-235-LABEL"`
					LABEL1 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-236-LABEL"`
					LABEL2 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-237-LABEL"`
					LABEL3 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-238-LABEL"`
					INPUT struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-239-INPUT"`
					INPUT1 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-240-INPUT"`
					INPUT2 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-241-INPUT"`
					INPUT3 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-242-INPUT"`
					INPUT4 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-243-INPUT"`
					INPUT5 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-244-INPUT"`
					EM struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-245-EM"`
					DIV struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-246-DIV"`
					BUTTON struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-247-BUTTON"`
					BUTTON1 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-248-BUTTON"`
					BUTTON2 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-249-BUTTON"`
					BUTTON3 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-250-BUTTON"`
					DIV1 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-251-DIV"`
					BUTTON4 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-252-BUTTON"`
					BUTTON5 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-253-BUTTON"`
					BUTTON6 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-254-BUTTON"`
					DIV2 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-255-DIV"`
					DIV3 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-256-DIV"`
					STRONG struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-257-STRONG"`
					DIV4 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-258-DIV"`
					DIV5 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-259-DIV"`
					DIV6 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-260-DIV"`
					DIV7 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-261-DIV"`
					DIV8 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-262-DIV"`
					DIV9 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-263-DIV"`
					H4 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-264-H4"`
					DIV10 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-265-DIV"`
					DIV11 struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-266-DIV"`
					VIDEO struct {
						Top    float64 `json:"top"`
						Bottom float64 `json:"bottom"`
						Left   float64 `json:"left"`
						Right  float64 `json:"right"`
						Width  float64 `json:"width"`
						Height float64 `json:"height"`
					} `json:"4-267-VIDEO"`
				} `json:"nodes"`
			} `json:"details"`
		} `json:"full-page-screenshot"`
		ScriptTreemapData struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type  string `json:"type"`
				Nodes []struct {
					Name          string  `json:"name"`
					ResourceBytes float64 `json:"resourceBytes"`
					Children      []struct {
						Name          string  `json:"name"`
						ResourceBytes float64 `json:"resourceBytes"`
						Children      []struct {
							Name          string  `json:"name"`
							ResourceBytes float64 `json:"resourceBytes"`
							Children      []struct {
								Name          string  `json:"name"`
								ResourceBytes float64 `json:"resourceBytes"`
								UnusedBytes   float64 `json:"unusedBytes,omitempty"`
								Children      []struct {
									Name                           string  `json:"name"`
									ResourceBytes                  float64 `json:"resourceBytes"`
									UnusedBytes                    float64 `json:"unusedBytes,omitempty"`
									DuplicatedNormalizedModuleName string  `json:"duplicatedNormalizedModuleName,omitempty"`
									Children                       []struct {
										Name          string  `json:"name"`
										ResourceBytes float64 `json:"resourceBytes"`
										UnusedBytes   float64 `json:"unusedBytes,omitempty"`
										Children      []struct {
											Name          string  `json:"name"`
											ResourceBytes float64 `json:"resourceBytes"`
											UnusedBytes   float64 `json:"unusedBytes"`
										} `json:"children,omitempty"`
									} `json:"children,omitempty"`
								} `json:"children,omitempty"`
							} `json:"children,omitempty"`
							UnusedBytes                    float64 `json:"unusedBytes,omitempty"`
							DuplicatedNormalizedModuleName string  `json:"duplicatedNormalizedModuleName,omitempty"`
						} `json:"children,omitempty"`
						UnusedBytes float64 `json:"unusedBytes,omitempty"`
					} `json:"children,omitempty"`
					UnusedBytes float64 `json:"unusedBytes"`
				} `json:"nodes"`
			} `json:"details"`
		} `json:"script-treemap-data"`
		PwaCrossBrowser struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"pwa-cross-browser"`
		PwaPageTransitions struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"pwa-page-transitions"`
		PwaEachPageHasUrl struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"pwa-each-page-has-url"`
		Accesskeys struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"accesskeys"`
		AriaAllowedAttr struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"aria-allowed-attr"`
		AriaCommandName struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"aria-command-name"`
		AriaHiddenBody struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"aria-hidden-body"`
		AriaHiddenFocus struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"aria-hidden-focus"`
		AriaInputFieldName struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"aria-input-field-name"`
		AriaMeterName struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"aria-meter-name"`
		AriaProgressbarName struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"aria-progressbar-name"`
		AriaRequiredAttr struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"aria-required-attr"`
		AriaRequiredChildren struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key             string `json:"key"`
					ItemType        string `json:"itemType"`
					SubItemsHeading struct {
						Key      string `json:"key"`
						ItemType string `json:"itemType"`
					} `json:"subItemsHeading"`
					Text string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet     string `json:"snippet"`
						NodeLabel   string `json:"nodeLabel"`
						Explanation string `json:"explanation"`
					} `json:"node"`
				} `json:"items"`
				DebugData struct {
					Type   string   `json:"type"`
					Impact string   `json:"impact"`
					Tags   []string `json:"tags"`
				} `json:"debugData"`
			} `json:"details"`
		} `json:"aria-required-children"`
		AriaRequiredParent struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key             string `json:"key"`
					ItemType        string `json:"itemType"`
					SubItemsHeading struct {
						Key      string `json:"key"`
						ItemType string `json:"itemType"`
					} `json:"subItemsHeading"`
					Text string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet     string `json:"snippet"`
						NodeLabel   string `json:"nodeLabel"`
						Explanation string `json:"explanation"`
					} `json:"node"`
				} `json:"items"`
				DebugData struct {
					Type   string   `json:"type"`
					Impact string   `json:"impact"`
					Tags   []string `json:"tags"`
				} `json:"debugData"`
			} `json:"details"`
		} `json:"aria-required-parent"`
		AriaRoles struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"aria-roles"`
		AriaToggleFieldName struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"aria-toggle-field-name"`
		AriaTooltipName struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"aria-tooltip-name"`
		AriaTreeitemName struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"aria-treeitem-name"`
		AriaValidAttrValue struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"aria-valid-attr-value"`
		AriaValidAttr struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"aria-valid-attr"`
		ButtonName struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"button-name"`
		Bypass struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"bypass"`
		ColorContrast struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key             string `json:"key"`
					ItemType        string `json:"itemType"`
					SubItemsHeading struct {
						Key      string `json:"key"`
						ItemType string `json:"itemType"`
					} `json:"subItemsHeading"`
					Text string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet     string `json:"snippet"`
						NodeLabel   string `json:"nodeLabel"`
						Explanation string `json:"explanation"`
					} `json:"node"`
					SubItems struct {
						Type  string `json:"type"`
						Items []struct {
							RelatedNode struct {
								Type         string `json:"type"`
								LhId         string `json:"lhId"`
								Path         string `json:"path"`
								Selector     string `json:"selector"`
								BoundingRect struct {
									Top    float64 `json:"top"`
									Bottom float64 `json:"bottom"`
									Left   float64 `json:"left"`
									Right  float64 `json:"right"`
									Width  float64 `json:"width"`
									Height float64 `json:"height"`
								} `json:"boundingRect"`
								Snippet   string `json:"snippet"`
								NodeLabel string `json:"nodeLabel"`
							} `json:"relatedNode"`
						} `json:"items"`
					} `json:"subItems"`
				} `json:"items"`
				DebugData struct {
					Type   string   `json:"type"`
					Impact string   `json:"impact"`
					Tags   []string `json:"tags"`
				} `json:"debugData"`
			} `json:"details"`
		} `json:"color-contrast"`
		DefinitionList struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"definition-list"`
		Dlitem struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"dlitem"`
		DocumentTitle struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"document-title"`
		DuplicateIdActive struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"duplicate-id-active"`
		DuplicateIdAria struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"duplicate-id-aria"`
		FormFieldMultipleLabels struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"form-field-multiple-labels"`
		FrameTitle struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"frame-title"`
		HeadingOrder struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key             string `json:"key"`
					ItemType        string `json:"itemType"`
					SubItemsHeading struct {
						Key      string `json:"key"`
						ItemType string `json:"itemType"`
					} `json:"subItemsHeading"`
					Text string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet     string `json:"snippet"`
						NodeLabel   string `json:"nodeLabel"`
						Explanation string `json:"explanation"`
					} `json:"node"`
				} `json:"items"`
				DebugData struct {
					Type   string   `json:"type"`
					Impact string   `json:"impact"`
					Tags   []string `json:"tags"`
				} `json:"debugData"`
			} `json:"details"`
		} `json:"heading-order"`
		HtmlHasLang struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"html-has-lang"`
		HtmlLangValid struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"html-lang-valid"`
		ImageAlt struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"image-alt"`
		InputImageAlt struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"input-image-alt"`
		Label struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"label"`
		LinkName struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"link-name"`
		List struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"list"`
		Listitem struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"listitem"`
		MetaRefresh struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"meta-refresh"`
		MetaViewport struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"meta-viewport"`
		ObjectAlt struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"object-alt"`
		Tabindex struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"tabindex"`
		TdHeadersAttr struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"td-headers-attr"`
		ThHasDataCells struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"th-has-data-cells"`
		ValidLang struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"valid-lang"`
		VideoCaption struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"video-caption"`
		CustomControlsLabels struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"custom-controls-labels"`
		CustomControlsRoles struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"custom-controls-roles"`
		FocusTraps struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"focus-traps"`
		FocusableControls struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"focusable-controls"`
		InteractiveElementAffordance struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"interactive-element-affordance"`
		LogicalTabOrder struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"logical-tab-order"`
		ManagedFocus struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"managed-focus"`
		OffscreenContentHidden struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"offscreen-content-hidden"`
		UseLandmarks struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"use-landmarks"`
		VisualOrderFollowsDom struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"visual-order-follows-dom"`
		UsesLongCacheTtl struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			NumericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key         string  `json:"key"`
					ItemType    string  `json:"itemType"`
					Text        string  `json:"text"`
					DisplayUnit string  `json:"displayUnit,omitempty"`
					Granularity float64 `json:"granularity,omitempty"`
				} `json:"headings"`
				Items []struct {
					Url                 string  `json:"url"`
					CacheLifetimeMs     float64 `json:"cacheLifetimeMs"`
					CacheHitProbability float64 `json:"cacheHitProbability"`
					TotalBytes          float64 `json:"totalBytes"`
					WastedBytes         float64 `json:"wastedBytes"`
					DebugData           struct {
						Type   string  `json:"type"`
						MaxAge float64 `json:"max-age"`
					} `json:"debugData,omitempty"`
				} `json:"items"`
				Summary struct {
					WastedBytes float64 `json:"wastedBytes"`
				} `json:"summary"`
			} `json:"details"`
		} `json:"uses-long-cache-ttl"`
		TotalByteWeight struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key      string `json:"key"`
					ItemType string `json:"itemType"`
					Text     string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Url        string  `json:"url"`
					TotalBytes float64 `json:"totalBytes"`
				} `json:"items"`
			} `json:"details"`
		} `json:"total-byte-weight"`
		OffscreenImages struct {
			Id               string        `json:"id"`
			Title            string        `json:"title"`
			Description      string        `json:"description"`
			score            float64       `json:"score"`
			ScoreDisplayMode string        `json:"scoreDisplayMode"`
			numericValue     float64       `json:"numericValue"`
			NumericUnit      string        `json:"numericUnit"`
			DisplayValue     string        `json:"displayValue"`
			Warnings         []interface{} `json:"warnings"`
			Details          struct {
				Type                string        `json:"type"`
				Headings            []interface{} `json:"headings"`
				Items               []interface{} `json:"items"`
				OverallSavingsMs    float64       `json:"overallSavingsMs"`
				OverallSavingsBytes float64       `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"offscreen-images"`
		RenderBlockingResources struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key       string `json:"key"`
					ValueType string `json:"valueType"`
					Label     string `json:"label"`
				} `json:"headings"`
				Items []struct {
					Url        string  `json:"url"`
					TotalBytes float64 `json:"totalBytes"`
					WastedMs   float64 `json:"wastedMs"`
				} `json:"items"`
				OverallSavingsMs float64 `json:"overallSavingsMs"`
			} `json:"details"`
		} `json:"render-blocking-resources"`
		UnminifiedCss struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type                string        `json:"type"`
				Headings            []interface{} `json:"headings"`
				Items               []interface{} `json:"items"`
				OverallSavingsMs    float64       `json:"overallSavingsMs"`
				OverallSavingsBytes float64       `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"unminified-css"`
		UnminifiedJavascript struct {
			Id               string        `json:"id"`
			Title            string        `json:"title"`
			Description      string        `json:"description"`
			score            float64       `json:"score"`
			ScoreDisplayMode string        `json:"scoreDisplayMode"`
			numericValue     float64       `json:"numericValue"`
			NumericUnit      string        `json:"numericUnit"`
			DisplayValue     string        `json:"displayValue"`
			Warnings         []interface{} `json:"warnings"`
			Details          struct {
				Type                string        `json:"type"`
				Headings            []interface{} `json:"headings"`
				Items               []interface{} `json:"items"`
				OverallSavingsMs    float64       `json:"overallSavingsMs"`
				OverallSavingsBytes float64       `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"unminified-javascript"`
		UnusedCssRules struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key       string `json:"key"`
					ValueType string `json:"valueType"`
					Label     string `json:"label"`
				} `json:"headings"`
				Items []struct {
					Url           string  `json:"url"`
					WastedBytes   float64 `json:"wastedBytes"`
					WastedPercent float64 `json:"wastedPercent"`
					TotalBytes    float64 `json:"totalBytes"`
				} `json:"items"`
				OverallSavingsMs    float64 `json:"overallSavingsMs"`
				OverallSavingsBytes float64 `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"unused-css-rules"`
		UnusedJavascript struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key             string `json:"key"`
					ValueType       string `json:"valueType"`
					SubItemsHeading struct {
						Key       string `json:"key"`
						ValueType string `json:"valueType,omitempty"`
					} `json:"subItemsHeading"`
					Label string `json:"label"`
				} `json:"headings"`
				Items []struct {
					Url           string  `json:"url"`
					TotalBytes    float64 `json:"totalBytes"`
					WastedBytes   float64 `json:"wastedBytes"`
					WastedPercent float64 `json:"wastedPercent"`
				} `json:"items"`
				OverallSavingsMs    float64 `json:"overallSavingsMs"`
				OverallSavingsBytes float64 `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"unused-javascript"`
		ModernImageFormats struct {
			Id               string        `json:"id"`
			Title            string        `json:"title"`
			Description      string        `json:"description"`
			score            float64       `json:"score"`
			ScoreDisplayMode string        `json:"scoreDisplayMode"`
			numericValue     float64       `json:"numericValue"`
			NumericUnit      string        `json:"numericUnit"`
			DisplayValue     string        `json:"displayValue"`
			Warnings         []interface{} `json:"warnings"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key       string `json:"key"`
					ValueType string `json:"valueType"`
					Label     string `json:"label"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet   string `json:"snippet"`
						NodeLabel string `json:"nodeLabel"`
					} `json:"node"`
					Url             string  `json:"url"`
					FromProtocol    bool    `json:"fromProtocol"`
					IsCrossOrigin   bool    `json:"isCrossOrigin"`
					TotalBytes      float64 `json:"totalBytes"`
					WastedBytes     float64 `json:"wastedBytes"`
					WastedWebpBytes float64 `json:"wastedWebpBytes"`
				} `json:"items"`
				OverallSavingsMs    float64 `json:"overallSavingsMs"`
				OverallSavingsBytes float64 `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"modern-image-formats"`
		UsesOptimizedImages struct {
			Id               string        `json:"id"`
			Title            string        `json:"title"`
			Description      string        `json:"description"`
			score            float64       `json:"score"`
			ScoreDisplayMode string        `json:"scoreDisplayMode"`
			numericValue     float64       `json:"numericValue"`
			NumericUnit      string        `json:"numericUnit"`
			DisplayValue     string        `json:"displayValue"`
			Warnings         []interface{} `json:"warnings"`
			Details          struct {
				Type                string        `json:"type"`
				Headings            []interface{} `json:"headings"`
				Items               []interface{} `json:"items"`
				OverallSavingsMs    float64       `json:"overallSavingsMs"`
				OverallSavingsBytes float64       `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"uses-optimized-images"`
		UsesTextCompression struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key       string `json:"key"`
					ValueType string `json:"valueType"`
					Label     string `json:"label"`
				} `json:"headings"`
				Items []struct {
					Url         string  `json:"url"`
					TotalBytes  float64 `json:"totalBytes"`
					WastedBytes float64 `json:"wastedBytes"`
				} `json:"items"`
				OverallSavingsMs    float64 `json:"overallSavingsMs"`
				OverallSavingsBytes float64 `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"uses-text-compression"`
		UsesResponsiveImages struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key       string `json:"key"`
					ValueType string `json:"valueType"`
					Label     string `json:"label"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet   string `json:"snippet"`
						NodeLabel string `json:"nodeLabel"`
					} `json:"node"`
					Url           string  `json:"url"`
					TotalBytes    float64 `json:"totalBytes"`
					WastedBytes   float64 `json:"wastedBytes"`
					WastedPercent float64 `json:"wastedPercent"`
				} `json:"items"`
				OverallSavingsMs    float64 `json:"overallSavingsMs"`
				OverallSavingsBytes float64 `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"uses-responsive-images"`
		EfficientAnimatedContent struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type                string        `json:"type"`
				Headings            []interface{} `json:"headings"`
				Items               []interface{} `json:"items"`
				OverallSavingsMs    float64       `json:"overallSavingsMs"`
				OverallSavingsBytes float64       `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"efficient-animated-content"`
		DuplicatedJavascript struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key             *string `json:"key"`
					ValueType       string  `json:"valueType"`
					SubItemsHeading struct {
						Key       string `json:"key"`
						ValueType string `json:"valueType,omitempty"`
					} `json:"subItemsHeading,omitempty"`
					Label       string  `json:"label"`
					Granularity float64 `json:"granularity,omitempty"`
				} `json:"headings"`
				Items []struct {
					Source      string  `json:"source"`
					WastedBytes float64 `json:"wastedBytes"`
					Url         string  `json:"url"`
					TotalBytes  float64 `json:"totalBytes"`
					SubItems    struct {
						Type  string `json:"type"`
						Items []struct {
							Url                 string  `json:"url"`
							SourceTransferBytes float64 `json:"sourceTransferBytes"`
						} `json:"items"`
					} `json:"subItems"`
				} `json:"items"`
				OverallSavingsMs    float64 `json:"overallSavingsMs"`
				OverallSavingsBytes float64 `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"duplicated-javascript"`
		LegacyJavascript struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type                string        `json:"type"`
				Headings            []interface{} `json:"headings"`
				Items               []interface{} `json:"items"`
				OverallSavingsMs    float64       `json:"overallSavingsMs"`
				OverallSavingsBytes float64       `json:"overallSavingsBytes"`
			} `json:"details"`
		} `json:"legacy-javascript"`
		Doctype struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
		} `json:"doctype"`
		Charset struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
		} `json:"charset"`
		DomSize struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			Score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key      string `json:"key"`
					ItemType string `json:"itemType"`
					Text     string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Statistic string  `json:"statistic"`
					Value     float64 `json:"value"`
					Node      struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet   string `json:"snippet"`
						NodeLabel string `json:"nodeLabel"`
					} `json:"node,omitempty"`
				} `json:"items"`
			} `json:"details"`
		} `json:"dom-size"`
		GeolocationOnStart struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"geolocation-on-start"`
		InspectorIssues struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"inspector-issues"`
		NoDocumentWrite struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"no-document-write"`
		NoVulnerableLibraries struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
				Summary  struct {
				} `json:"summary"`
			} `json:"details"`
		} `json:"no-vulnerable-libraries"`
		JsLibraries struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key      string `json:"key"`
					ItemType string `json:"itemType"`
					Text     string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Name    string `json:"name"`
					Version string `json:"version"`
					Npm     string `json:"npm"`
				} `json:"items"`
				Summary struct {
				} `json:"summary"`
				DebugData struct {
					Type   string `json:"type"`
					Stacks []struct {
						Id      string `json:"id"`
						Version string `json:"version"`
					} `json:"stacks"`
				} `json:"debugData"`
			} `json:"details"`
		} `json:"js-libraries"`
		NotificationOnStart struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"notification-on-start"`
		PasswordInputsCanBePastedInto struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"password-inputs-can-be-pasted-into"`
		UsesHttp2 struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			numericValue     float64 `json:"numericValue"`
			NumericUnit      string  `json:"numericUnit"`
			Details          struct {
				Type             string        `json:"type"`
				Headings         []interface{} `json:"headings"`
				Items            []interface{} `json:"items"`
				OverallSavingsMs float64       `json:"overallSavingsMs"`
			} `json:"details"`
		} `json:"uses-http2"`
		UsesPassiveEventListeners struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key      string `json:"key"`
					ItemType string `json:"itemType"`
					Text     string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Source struct {
						Type        string  `json:"type"`
						Url         string  `json:"url"`
						UrlProvider string  `json:"urlProvider"`
						Line        float64 `json:"line"`
						Column      float64 `json:"column"`
					} `json:"source"`
				} `json:"items"`
			} `json:"details"`
		} `json:"uses-passive-event-listeners"`
		MetaDescription struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
		} `json:"meta-description"`
		HttpStatusCode struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
		} `json:"http-status-code"`
		FontSize struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key      string `json:"key"`
					ItemType string `json:"itemType"`
					Text     string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Source struct {
						Type  string `json:"type"`
						Value string `json:"value"`
					} `json:"source"`
					Selector string `json:"selector"`
					Coverage string `json:"coverage"`
					FontSize string `json:"fontSize"`
				} `json:"items"`
			} `json:"details"`
		} `json:"font-size"`
		LinkText struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
				Summary  struct {
				} `json:"summary"`
			} `json:"details"`
		} `json:"link-text"`
		CrawlableAnchors struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string `json:"type"`
				Headings []struct {
					Key      string `json:"key"`
					ItemType string `json:"itemType"`
					Text     string `json:"text"`
				} `json:"headings"`
				Items []struct {
					Node struct {
						Type         string `json:"type"`
						LhId         string `json:"lhId"`
						Path         string `json:"path"`
						Selector     string `json:"selector"`
						BoundingRect struct {
							Top    float64 `json:"top"`
							Bottom float64 `json:"bottom"`
							Left   float64 `json:"left"`
							Right  float64 `json:"right"`
							Width  float64 `json:"width"`
							Height float64 `json:"height"`
						} `json:"boundingRect"`
						Snippet   string `json:"snippet"`
						NodeLabel string `json:"nodeLabel"`
					} `json:"node"`
				} `json:"items"`
			} `json:"details"`
		} `json:"crawlable-anchors"`
		IsCrawlable struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"is-crawlable"`
		RobotsTxt struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
				Summary  struct {
				} `json:"summary"`
			} `json:"details"`
		} `json:"robots-txt"`
		TapTargets struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			DisplayValue     string  `json:"displayValue"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"tap-targets"`
		Hreflang struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"hreflang"`
		Plugins struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
			Details          struct {
				Type     string        `json:"type"`
				Headings []interface{} `json:"headings"`
				Items    []interface{} `json:"items"`
			} `json:"details"`
		} `json:"plugins"`
		Canonical struct {
			Id               string  `json:"id"`
			Title            string  `json:"title"`
			Description      string  `json:"description"`
			score            float64 `json:"score"`
			ScoreDisplayMode string  `json:"scoreDisplayMode"`
		} `json:"canonical"`
		StructuredData struct {
			Id               string      `json:"id"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Score            interface{} `json:"score"`
			ScoreDisplayMode string      `json:"scoreDisplayMode"`
		} `json:"structured-data"`
	} `json:"audits"`
	ConfigSettings struct {
		Output         string  `json:"output"`
		MaxWaitForFcp  float64 `json:"maxWaitForFcp"`
		MaxWaitForLoad float64 `json:"maxWaitForLoad"`
		FormFactor     string  `json:"formFactor"`
		Throttling     struct {
			RttMs                  float64 `json:"rttMs"`
			ThroughputKbps         float64 `json:"throughputKbps"`
			RequestLatencyMs       float64 `json:"requestLatencyMs"`
			DownloadThroughputKbps float64 `json:"downloadThroughputKbps"`
			UploadThroughputKbps   float64 `json:"uploadThroughputKbps"`
			CpuSlowdownMultiplier  float64 `json:"cpuSlowdownMultiplier"`
		} `json:"throttling"`
		ThrottlingMethod string `json:"throttlingMethod"`
		ScreenEmulation  struct {
			Mobile            bool    `json:"mobile"`
			Width             float64 `json:"width"`
			Height            float64 `json:"height"`
			DeviceScaleFactor float64 `json:"deviceScaleFactor"`
			Disabled          bool    `json:"disabled"`
		} `json:"screenEmulation"`
		EmulatedUserAgent         string      `json:"emulatedUserAgent"`
		AuditMode                 bool        `json:"auditMode"`
		GatherMode                bool        `json:"gatherMode"`
		DisableStorageReset       bool        `json:"disableStorageReset"`
		DebugNavigation           bool        `json:"debugNavigation"`
		Channel                   string      `json:"channel"`
		Budgets                   interface{} `json:"budgets"`
		Locale                    string      `json:"locale"`
		BlockedUrlPatterns        interface{} `json:"blockedUrlPatterns"`
		AdditionalTraceCategories interface{} `json:"additionalTraceCategories"`
		ExtraHeaders              interface{} `json:"extraHeaders"`
		PrecomputedLanternData    interface{} `json:"precomputedLanternData"`
		OnlyAudits                interface{} `json:"onlyAudits"`
		OnlyCategories            interface{} `json:"onlyCategories"`
		SkipAudits                interface{} `json:"skipAudits"`
	} `json:"configSettings"`
	Categories struct {
		Performance struct {
			Title          string   `json:"title"`
			SupportedModes []string `json:"supportedModes"`
			AuditRefs      []struct {
				Id             string   `json:"id"`
				Weight         float64  `json:"weight"`
				Group          string   `json:"group,omitempty"`
				Acronym        string   `json:"acronym,omitempty"`
				RelevantAudits []string `json:"relevantAudits,omitempty"`
			} `json:"auditRefs"`
			Id    string  `json:"id"`
			Score float64 `json:"score"`
		} `json:"performance"`
		Accessibility struct {
			Title             string   `json:"title"`
			Description       string   `json:"description"`
			ManualDescription string   `json:"manualDescription"`
			SupportedModes    []string `json:"supportedModes"`
			AuditRefs         []struct {
				Id     string  `json:"id"`
				Weight float64 `json:"weight"`
				Group  string  `json:"group,omitempty"`
			} `json:"auditRefs"`
			Id    string  `json:"id"`
			Score float64 `json:"score"`
		} `json:"accessibility"`
		BestPractices struct {
			Title          string   `json:"title"`
			SupportedModes []string `json:"supportedModes"`
			AuditRefs      []struct {
				Id     string  `json:"id"`
				Weight float64 `json:"weight"`
				Group  string  `json:"group"`
			} `json:"auditRefs"`
			Id    string  `json:"id"`
			Score float64 `json:"score"`
		} `json:"best-practices"`
		Seo struct {
			Title             string   `json:"title"`
			Description       string   `json:"description"`
			ManualDescription string   `json:"manualDescription"`
			SupportedModes    []string `json:"supportedModes"`
			AuditRefs         []struct {
				Id     string  `json:"id"`
				Weight float64 `json:"weight"`
				Group  string  `json:"group,omitempty"`
			} `json:"auditRefs"`
			Id    string  `json:"id"`
			Score float64 `json:"score"`
		} `json:"seo"`
		Pwa struct {
			Title             string   `json:"title"`
			Description       string   `json:"description"`
			ManualDescription string   `json:"manualDescription"`
			SupportedModes    []string `json:"supportedModes"`
			AuditRefs         []struct {
				Id     string  `json:"id"`
				Weight float64 `json:"weight"`
				Group  string  `json:"group,omitempty"`
			} `json:"auditRefs"`
			Id    string  `json:"id"`
			Score float64 `json:"score"`
		} `json:"pwa"`
	} `json:"categories"`
	CategoryGroups struct {
		Metrics struct {
			Title string `json:"title"`
		} `json:"metrics"`
		LoadOpportunities struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"load-opportunities"`
		Budgets struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"budgets"`
		Diagnostics struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"diagnostics"`
		PwaInstallable struct {
			Title string `json:"title"`
		} `json:"pwa-installable"`
		PwaOptimized struct {
			Title string `json:"title"`
		} `json:"pwa-optimized"`
		A11YBestPractices struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"a11y-best-practices"`
		A11YColorContrast struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"a11y-color-contrast"`
		A11YNamesLabels struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"a11y-names-labels"`
		A11YNavigation struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"a11y-navigation"`
		A11YAria struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"a11y-aria"`
		A11YLanguage struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"a11y-language"`
		A11YAudioVideo struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"a11y-audio-video"`
		A11YTablesLists struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"a11y-tables-lists"`
		SeoMobile struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"seo-mobile"`
		SeoContent struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"seo-content"`
		SeoCrawl struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"seo-crawl"`
		BestPracticesTrustSafety struct {
			Title string `json:"title"`
		} `json:"best-practices-trust-safety"`
		BestPracticesUx struct {
			Title string `json:"title"`
		} `json:"best-practices-ux"`
		BestPracticesBrowserCompat struct {
			Title string `json:"title"`
		} `json:"best-practices-browser-compat"`
		BestPracticesGeneral struct {
			Title string `json:"title"`
		} `json:"best-practices-general"`
		Hidden struct {
			Title string `json:"title"`
		} `json:"hidden"`
	} `json:"categoryGroups"`
	StackPacks []interface{} `json:"stackPacks"`
	Timing     struct {
		Entries []struct {
			StartTime float64 `json:"startTime"`
			Name      string  `json:"name"`
			Duration  float64 `json:"duration"`
			EntryType string  `json:"entryType"`
		} `json:"entries"`
		Total float64 `json:"total"`
	} `json:"timing"`
	I18N struct {
		RendererFormattedStrings struct {
			CalculatorLink                   string `json:"calculatorLink"`
			CollapseView                     string `json:"collapseView"`
			CrcInitialNavigation             string `json:"crcInitialNavigation"`
			CrcLongestDurationLabel          string `json:"crcLongestDurationLabel"`
			DropdownCopyJSON                 string `json:"dropdownCopyJSON"`
			DropdownDarkTheme                string `json:"dropdownDarkTheme"`
			DropdownPrintExpanded            string `json:"dropdownPrintExpanded"`
			DropdownPrintSummary             string `json:"dropdownPrintSummary"`
			DropdownSaveGist                 string `json:"dropdownSaveGist"`
			DropdownSaveHTML                 string `json:"dropdownSaveHTML"`
			DropdownSaveJSON                 string `json:"dropdownSaveJSON"`
			DropdownViewer                   string `json:"dropdownViewer"`
			ErrorLabel                       string `json:"errorLabel"`
			ErrorMissingAuditInfo            string `json:"errorMissingAuditInfo"`
			ExpandView                       string `json:"expandView"`
			FooterIssue                      string `json:"footerIssue"`
			Hide                             string `json:"hide"`
			LabDataTitle                     string `json:"labDataTitle"`
			LsPerformanceCategoryDescription string `json:"lsPerformanceCategoryDescription"`
			ManualAuditsGroupTitle           string `json:"manualAuditsGroupTitle"`
			NotApplicableAuditsGroupTitle    string `json:"notApplicableAuditsGroupTitle"`
			OpportunityResourceColumnLabel   string `json:"opportunityResourceColumnLabel"`
			OpportunitySavingsColumnLabel    string `json:"opportunitySavingsColumnLabel"`
			PassedAuditsGroupTitle           string `json:"passedAuditsGroupTitle"`
			RuntimeAnalysisWindow            string `json:"runtimeAnalysisWindow"`
			RuntimeCustom                    string `json:"runtimeCustom"`
			RuntimeDesktopEmulation          string `json:"runtimeDesktopEmulation"`
			RuntimeMobileEmulation           string `json:"runtimeMobileEmulation"`
			RuntimeNoEmulation               string `json:"runtimeNoEmulation"`
			RuntimeSettingsAxeVersion        string `json:"runtimeSettingsAxeVersion"`
			RuntimeSettingsBenchmark         string `json:"runtimeSettingsBenchmark"`
			RuntimeSettingsCPUThrottling     string `json:"runtimeSettingsCPUThrottling"`
			RuntimeSettingsDevice            string `json:"runtimeSettingsDevice"`
			RuntimeSettingsNetworkThrottling string `json:"runtimeSettingsNetworkThrottling"`
			RuntimeSettingsUANetwork         string `json:"runtimeSettingsUANetwork"`
			RuntimeSingleLoad                string `json:"runtimeSingleLoad"`
			RuntimeSingleLoadTooltip         string `json:"runtimeSingleLoadTooltip"`
			RuntimeSlow4G                    string `json:"runtimeSlow4g"`
			RuntimeUnknown                   string `json:"runtimeUnknown"`
			Show                             string `json:"show"`
			ShowRelevantAudits               string `json:"showRelevantAudits"`
			SnippetCollapseButtonLabel       string `json:"snippetCollapseButtonLabel"`
			SnippetExpandButtonLabel         string `json:"snippetExpandButtonLabel"`
			ThirdPartyResourcesLabel         string `json:"thirdPartyResourcesLabel"`
			ThrottlingProvided               string `json:"throttlingProvided"`
			ToplevelWarningsMessage          string `json:"toplevelWarningsMessage"`
			VarianceDisclaimer               string `json:"varianceDisclaimer"`
			ViewOriginalTraceLabel           string `json:"viewOriginalTraceLabel"`
			ViewTraceLabel                   string `json:"viewTraceLabel"`
			ViewTreemapLabel                 string `json:"viewTreemapLabel"`
			WarningAuditsGroupTitle          string `json:"warningAuditsGroupTitle"`
			WarningHeader                    string `json:"warningHeader"`
		} `json:"rendererFormattedStrings"`
		IcuMessagePaths struct {
			LighthouseCoreGatherDriverNavigationJsWarningTimeout         []string `json:"lighthouse-core/gather/driver/navigation.js | warningTimeout"`
			LighthouseCoreAuditsIsOnHttpsJsTitle                         []string `json:"lighthouse-core/audits/is-on-https.js | title"`
			LighthouseCoreAuditsIsOnHttpsJsDescription                   []string `json:"lighthouse-core/audits/is-on-https.js | description"`
			LighthouseCoreAuditsServiceWorkerJsFailureTitle              []string `json:"lighthouse-core/audits/service-worker.js | failureTitle"`
			LighthouseCoreAuditsServiceWorkerJsDescription               []string `json:"lighthouse-core/audits/service-worker.js | description"`
			LighthouseCoreAuditsViewportJsTitle                          []string `json:"lighthouse-core/audits/viewport.js | title"`
			LighthouseCoreAuditsViewportJsDescription                    []string `json:"lighthouse-core/audits/viewport.js | description"`
			LighthouseCoreLibI18NI18NJsFirstContentfulPaintMetric        []string `json:"lighthouse-core/lib/i18n/i18n.js | firstContentfulPaintMetric"`
			LighthouseCoreAuditsMetricsFirstContentfulPaintJsDescription []string `json:"lighthouse-core/audits/metrics/first-contentful-paint.js | description"`
			LighthouseCoreLibI18NI18NJsSeconds                           []struct {
				Values struct {
					TimeInMs float64 `json:"timeInMs"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/lib/i18n/i18n.js | seconds"`
			LighthouseCoreLibI18NI18NJsLargestContentfulPaintMetric        []string `json:"lighthouse-core/lib/i18n/i18n.js | largestContentfulPaintMetric"`
			LighthouseCoreAuditsMetricsLargestContentfulPaintJsDescription []string `json:"lighthouse-core/audits/metrics/largest-contentful-paint.js | description"`
			LighthouseCoreLibI18NI18NJsFirstMeaningfulPaintMetric          []string `json:"lighthouse-core/lib/i18n/i18n.js | firstMeaningfulPaintMetric"`
			LighthouseCoreAuditsMetricsFirstMeaningfulPaintJsDescription   []string `json:"lighthouse-core/audits/metrics/first-meaningful-paint.js | description"`
			LighthouseCoreLibI18NI18NJsSpeedIndexMetric                    []string `json:"lighthouse-core/lib/i18n/i18n.js | speedIndexMetric"`
			LighthouseCoreAuditsMetricsSpeedIndexJsDescription             []string `json:"lighthouse-core/audits/metrics/speed-index.js | description"`
			LighthouseCoreLibI18NI18NJsTotalBlockingTimeMetric             []string `json:"lighthouse-core/lib/i18n/i18n.js | totalBlockingTimeMetric"`
			LighthouseCoreAuditsMetricsTotalBlockingTimeJsDescription      []string `json:"lighthouse-core/audits/metrics/total-blocking-time.js | description"`
			LighthouseCoreLibI18NI18NJsMs                                  []struct {
				Values struct {
					TimeInMs float64 `json:"timeInMs"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/lib/i18n/i18n.js | ms"`
			LighthouseCoreLibI18NI18NJsMaxPotentialFIDMetric              []string `json:"lighthouse-core/lib/i18n/i18n.js | maxPotentialFIDMetric"`
			LighthouseCoreAuditsMetricsMaxPotentialFidJsDescription       []string `json:"lighthouse-core/audits/metrics/max-potential-fid.js | description"`
			LighthouseCoreLibI18NI18NJsCumulativeLayoutShiftMetric        []string `json:"lighthouse-core/lib/i18n/i18n.js | cumulativeLayoutShiftMetric"`
			LighthouseCoreAuditsMetricsCumulativeLayoutShiftJsDescription []string `json:"lighthouse-core/audits/metrics/cumulative-layout-shift.js | description"`
			LighthouseCoreAuditsErrorsInConsoleJsTitle                    []string `json:"lighthouse-core/audits/errors-in-console.js | title"`
			LighthouseCoreAuditsErrorsInConsoleJsDescription              []string `json:"lighthouse-core/audits/errors-in-console.js | description"`
			LighthouseCoreAuditsServerResponseTimeJsTitle                 []string `json:"lighthouse-core/audits/server-response-time.js | title"`
			LighthouseCoreAuditsServerResponseTimeJsDescription           []string `json:"lighthouse-core/audits/server-response-time.js | description"`
			LighthouseCoreAuditsServerResponseTimeJsDisplayValue          []struct {
				Values struct {
					TimeInMs float64 `json:"timeInMs"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/audits/server-response-time.js | displayValue"`
			LighthouseCoreLibI18NI18NJsColumnURL                    []string `json:"lighthouse-core/lib/i18n/i18n.js | columnURL"`
			LighthouseCoreLibI18NI18NJsColumnTimeSpent              []string `json:"lighthouse-core/lib/i18n/i18n.js | columnTimeSpent"`
			LighthouseCoreLibI18NI18NJsInteractiveMetric            []string `json:"lighthouse-core/lib/i18n/i18n.js | interactiveMetric"`
			LighthouseCoreAuditsMetricsInteractiveJsDescription     []string `json:"lighthouse-core/audits/metrics/interactive.js | description"`
			LighthouseCoreAuditsUserTimingsJsTitle                  []string `json:"lighthouse-core/audits/user-timings.js | title"`
			LighthouseCoreAuditsUserTimingsJsDescription            []string `json:"lighthouse-core/audits/user-timings.js | description"`
			LighthouseCoreAuditsCriticalRequestChainsJsTitle        []string `json:"lighthouse-core/audits/critical-request-chains.js | title"`
			LighthouseCoreAuditsCriticalRequestChainsJsDescription  []string `json:"lighthouse-core/audits/critical-request-chains.js | description"`
			LighthouseCoreAuditsCriticalRequestChainsJsDisplayValue []struct {
				Values struct {
					ItemCount float64 `json:"itemCount"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/audits/critical-request-chains.js | displayValue"`
			LighthouseCoreAuditsRedirectsJsTitle                 []string `json:"lighthouse-core/audits/redirects.js | title"`
			LighthouseCoreAuditsRedirectsJsDescription           []string `json:"lighthouse-core/audits/redirects.js | description"`
			LighthouseCoreAuditsInstallableManifestJsTitle       []string `json:"lighthouse-core/audits/installable-manifest.js | title"`
			LighthouseCoreAuditsInstallableManifestJsDescription []string `json:"lighthouse-core/audits/installable-manifest.js | description"`
			LighthouseCoreAuditsAppleTouchIconJsTitle            []string `json:"lighthouse-core/audits/apple-touch-icon.js | title"`
			LighthouseCoreAuditsAppleTouchIconJsDescription      []string `json:"lighthouse-core/audits/apple-touch-icon.js | description"`
			LighthouseCoreAuditsSplashScreenJsFailureTitle       []string `json:"lighthouse-core/audits/splash-screen.js | failureTitle"`
			LighthouseCoreAuditsSplashScreenJsDescription        []string `json:"lighthouse-core/audits/splash-screen.js | description"`
			LighthouseCoreAuditsThemedOmniboxJsFailureTitle      []string `json:"lighthouse-core/audits/themed-omnibox.js | failureTitle"`
			LighthouseCoreAuditsThemedOmniboxJsDescription       []string `json:"lighthouse-core/audits/themed-omnibox.js | description"`
			LighthouseCoreAuditsMaskableIconJsFailureTitle       []string `json:"lighthouse-core/audits/maskable-icon.js | failureTitle"`
			LighthouseCoreAuditsMaskableIconJsDescription        []string `json:"lighthouse-core/audits/maskable-icon.js | description"`
			LighthouseCoreAuditsContentWidthJsTitle              []string `json:"lighthouse-core/audits/content-width.js | title"`
			LighthouseCoreAuditsContentWidthJsDescription        []string `json:"lighthouse-core/audits/content-width.js | description"`
			LighthouseCoreAuditsImageAspectRatioJsTitle          []string `json:"lighthouse-core/audits/image-aspect-ratio.js | title"`
			LighthouseCoreAuditsImageAspectRatioJsDescription    []string `json:"lighthouse-core/audits/image-aspect-ratio.js | description"`
			LighthouseCoreAuditsImageSizeResponsiveJsTitle       []string `json:"lighthouse-core/audits/image-size-responsive.js | title"`
			LighthouseCoreAuditsImageSizeResponsiveJsDescription []string `json:"lighthouse-core/audits/image-size-responsive.js | description"`
			LighthouseCoreAuditsPreloadFontsJsTitle              []string `json:"lighthouse-core/audits/preload-fonts.js | title"`
			LighthouseCoreAuditsPreloadFontsJsDescription        []string `json:"lighthouse-core/audits/preload-fonts.js | description"`
			LighthouseCoreAuditsDeprecationsJsFailureTitle       []string `json:"lighthouse-core/audits/deprecations.js | failureTitle"`
			LighthouseCoreAuditsDeprecationsJsDescription        []string `json:"lighthouse-core/audits/deprecations.js | description"`
			LighthouseCoreAuditsDeprecationsJsDisplayValue       []struct {
				Values struct {
					ItemCount float64 `json:"itemCount"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/audits/deprecations.js | displayValue"`
			LighthouseCoreAuditsDeprecationsJsColumnDeprecate           []string `json:"lighthouse-core/audits/deprecations.js | columnDeprecate"`
			LighthouseCoreLibI18NI18NJsColumnSource                     []string `json:"lighthouse-core/lib/i18n/i18n.js | columnSource"`
			LighthouseCoreAuditsDeprecationsJsUnknownDeprecation        []string `json:"lighthouse-core/audits/deprecations.js | unknownDeprecation"`
			LighthouseCoreAuditsMainthreadWorkBreakdownJsFailureTitle   []string `json:"lighthouse-core/audits/mainthread-work-breakdown.js | failureTitle"`
			LighthouseCoreAuditsMainthreadWorkBreakdownJsDescription    []string `json:"lighthouse-core/audits/mainthread-work-breakdown.js | description"`
			LighthouseCoreAuditsMainthreadWorkBreakdownJsColumnCategory []string `json:"lighthouse-core/audits/mainthread-work-breakdown.js | columnCategory"`
			LighthouseCoreAuditsBootupTimeJsFailureTitle                []string `json:"lighthouse-core/audits/bootup-time.js | failureTitle"`
			LighthouseCoreAuditsBootupTimeJsDescription                 []string `json:"lighthouse-core/audits/bootup-time.js | description"`
			LighthouseCoreAuditsBootupTimeJsColumnTotal                 []string `json:"lighthouse-core/audits/bootup-time.js | columnTotal"`
			LighthouseCoreAuditsBootupTimeJsColumnScriptEval            []string `json:"lighthouse-core/audits/bootup-time.js | columnScriptEval"`
			LighthouseCoreAuditsBootupTimeJsColumnScriptParse           []string `json:"lighthouse-core/audits/bootup-time.js | columnScriptParse"`
			LighthouseCoreAuditsUsesRelPreloadJsTitle                   []string `json:"lighthouse-core/audits/uses-rel-preload.js | title"`
			LighthouseCoreAuditsUsesRelPreloadJsDescription             []string `json:"lighthouse-core/audits/uses-rel-preload.js | description"`
			LighthouseCoreAuditsUsesRelPreconnectJsTitle                []string `json:"lighthouse-core/audits/uses-rel-preconnect.js | title"`
			LighthouseCoreAuditsUsesRelPreconnectJsDescription          []string `json:"lighthouse-core/audits/uses-rel-preconnect.js | description"`
			LighthouseCoreAuditsFontDisplayJsTitle                      []string `json:"lighthouse-core/audits/font-display.js | title"`
			LighthouseCoreAuditsFontDisplayJsDescription                []string `json:"lighthouse-core/audits/font-display.js | description"`
			LighthouseCoreAuditsNetworkRttJsTitle                       []string `json:"lighthouse-core/audits/network-rtt.js | title"`
			LighthouseCoreAuditsNetworkRttJsDescription                 []string `json:"lighthouse-core/audits/network-rtt.js | description"`
			LighthouseCoreAuditsNetworkServerLatencyJsTitle             []string `json:"lighthouse-core/audits/network-server-latency.js | title"`
			LighthouseCoreAuditsNetworkServerLatencyJsDescription       []string `json:"lighthouse-core/audits/network-server-latency.js | description"`
			LighthouseCoreAuditsPerformanceBudgetJsTitle                []string `json:"lighthouse-core/audits/performance-budget.js | title"`
			LighthouseCoreAuditsPerformanceBudgetJsDescription          []string `json:"lighthouse-core/audits/performance-budget.js | description"`
			LighthouseCoreAuditsTimingBudgetJsTitle                     []string `json:"lighthouse-core/audits/timing-budget.js | title"`
			LighthouseCoreAuditsTimingBudgetJsDescription               []string `json:"lighthouse-core/audits/timing-budget.js | description"`
			LighthouseCoreAuditsResourceSummaryJsTitle                  []string `json:"lighthouse-core/audits/resource-summary.js | title"`
			LighthouseCoreAuditsResourceSummaryJsDescription            []string `json:"lighthouse-core/audits/resource-summary.js | description"`
			LighthouseCoreAuditsResourceSummaryJsDisplayValue           []struct {
				Values struct {
					RequestCount float64 `json:"requestCount"`
					ByteCount    float64 `json:"byteCount"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/audits/resource-summary.js | displayValue"`
			LighthouseCoreLibI18NI18NJsColumnResourceType                  []string `json:"lighthouse-core/lib/i18n/i18n.js | columnResourceType"`
			LighthouseCoreLibI18NI18NJsColumnRequests                      []string `json:"lighthouse-core/lib/i18n/i18n.js | columnRequests"`
			LighthouseCoreLibI18NI18NJsColumnTransferSize                  []string `json:"lighthouse-core/lib/i18n/i18n.js | columnTransferSize"`
			LighthouseCoreLibI18NI18NJsTotalResourceType                   []string `json:"lighthouse-core/lib/i18n/i18n.js | totalResourceType"`
			LighthouseCoreLibI18NI18NJsScriptResourceType                  []string `json:"lighthouse-core/lib/i18n/i18n.js | scriptResourceType"`
			LighthouseCoreLibI18NI18NJsImageResourceType                   []string `json:"lighthouse-core/lib/i18n/i18n.js | imageResourceType"`
			LighthouseCoreLibI18NI18NJsStylesheetResourceType              []string `json:"lighthouse-core/lib/i18n/i18n.js | stylesheetResourceType"`
			LighthouseCoreLibI18NI18NJsFontResourceType                    []string `json:"lighthouse-core/lib/i18n/i18n.js | fontResourceType"`
			LighthouseCoreLibI18NI18NJsDocumentResourceType                []string `json:"lighthouse-core/lib/i18n/i18n.js | documentResourceType"`
			LighthouseCoreLibI18NI18NJsOtherResourceType                   []string `json:"lighthouse-core/lib/i18n/i18n.js | otherResourceType"`
			LighthouseCoreLibI18NI18NJsMediaResourceType                   []string `json:"lighthouse-core/lib/i18n/i18n.js | mediaResourceType"`
			LighthouseCoreLibI18NI18NJsThirdPartyResourceType              []string `json:"lighthouse-core/lib/i18n/i18n.js | thirdPartyResourceType"`
			LighthouseCoreAuditsThirdPartySummaryJsTitle                   []string `json:"lighthouse-core/audits/third-party-summary.js | title"`
			LighthouseCoreAuditsThirdPartySummaryJsDescription             []string `json:"lighthouse-core/audits/third-party-summary.js | description"`
			LighthouseCoreAuditsThirdPartyFacadesJsTitle                   []string `json:"lighthouse-core/audits/third-party-facades.js | title"`
			LighthouseCoreAuditsThirdPartyFacadesJsDescription             []string `json:"lighthouse-core/audits/third-party-facades.js | description"`
			LighthouseCoreAuditsLargestContentfulPaintElementJsTitle       []string `json:"lighthouse-core/audits/largest-contentful-paint-element.js | title"`
			LighthouseCoreAuditsLargestContentfulPaintElementJsDescription []string `json:"lighthouse-core/audits/largest-contentful-paint-element.js | description"`
			LighthouseCoreLibI18NI18NJsDisplayValueElementsFound           []struct {
				Values struct {
					NodeCount float64 `json:"nodeCount"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/lib/i18n/i18n.js | displayValueElementsFound"`
			LighthouseCoreLibI18NI18NJsColumnElement                    []string `json:"lighthouse-core/lib/i18n/i18n.js | columnElement"`
			LighthouseCoreAuditsLcpLazyLoadedJsTitle                    []string `json:"lighthouse-core/audits/lcp-lazy-loaded.js | title"`
			LighthouseCoreAuditsLcpLazyLoadedJsDescription              []string `json:"lighthouse-core/audits/lcp-lazy-loaded.js | description"`
			LighthouseCoreAuditsLayoutShiftElementsJsTitle              []string `json:"lighthouse-core/audits/layout-shift-elements.js | title"`
			LighthouseCoreAuditsLayoutShiftElementsJsDescription        []string `json:"lighthouse-core/audits/layout-shift-elements.js | description"`
			LighthouseCoreAuditsLayoutShiftElementsJsColumnContribution []string `json:"lighthouse-core/audits/layout-shift-elements.js | columnContribution"`
			LighthouseCoreAuditsLongTasksJsTitle                        []string `json:"lighthouse-core/audits/long-tasks.js | title"`
			LighthouseCoreAuditsLongTasksJsDescription                  []string `json:"lighthouse-core/audits/long-tasks.js | description"`
			LighthouseCoreAuditsLongTasksJsDisplayValue                 []struct {
				Values struct {
					ItemCount float64 `json:"itemCount"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/audits/long-tasks.js | displayValue"`
			LighthouseCoreLibI18NI18NJsColumnStartTime                            []string `json:"lighthouse-core/lib/i18n/i18n.js | columnStartTime"`
			LighthouseCoreLibI18NI18NJsColumnDuration                             []string `json:"lighthouse-core/lib/i18n/i18n.js | columnDuration"`
			LighthouseCoreAuditsNoUnloadListenersJsTitle                          []string `json:"lighthouse-core/audits/no-unload-listeners.js | title"`
			LighthouseCoreAuditsNoUnloadListenersJsDescription                    []string `json:"lighthouse-core/audits/no-unload-listeners.js | description"`
			LighthouseCoreAuditsNonCompositedAnimationsJsTitle                    []string `json:"lighthouse-core/audits/non-composited-animations.js | title"`
			LighthouseCoreAuditsNonCompositedAnimationsJsDescription              []string `json:"lighthouse-core/audits/non-composited-animations.js | description"`
			LighthouseCoreAuditsUnsizedImagesJsTitle                              []string `json:"lighthouse-core/audits/unsized-images.js | title"`
			LighthouseCoreAuditsUnsizedImagesJsDescription                        []string `json:"lighthouse-core/audits/unsized-images.js | description"`
			LighthouseCoreAuditsValidSourceMapsJsFailureTitle                     []string `json:"lighthouse-core/audits/valid-source-maps.js | failureTitle"`
			LighthouseCoreAuditsValidSourceMapsJsDescription                      []string `json:"lighthouse-core/audits/valid-source-maps.js | description"`
			LighthouseCoreAuditsValidSourceMapsJsColumnMapURL                     []string `json:"lighthouse-core/audits/valid-source-maps.js | columnMapURL"`
			LighthouseCoreAuditsValidSourceMapsJsMissingSourceMapErrorMessage     []string `json:"lighthouse-core/audits/valid-source-maps.js | missingSourceMapErrorMessage"`
			LighthouseCoreAuditsPreloadLcpImageJsTitle                            []string `json:"lighthouse-core/audits/preload-lcp-image.js | title"`
			LighthouseCoreAuditsPreloadLcpImageJsDescription                      []string `json:"lighthouse-core/audits/preload-lcp-image.js | description"`
			LighthouseCoreLibI18NI18NJsColumnWastedBytes                          []string `json:"lighthouse-core/lib/i18n/i18n.js | columnWastedBytes"`
			LighthouseCoreAuditsCspXssJsTitle                                     []string `json:"lighthouse-core/audits/csp-xss.js | title"`
			LighthouseCoreAuditsCspXssJsDescription                               []string `json:"lighthouse-core/audits/csp-xss.js | description"`
			LighthouseCoreLibI18NI18NJsColumnDescription                          []string `json:"lighthouse-core/lib/i18n/i18n.js | columnDescription"`
			LighthouseCoreAuditsCspXssJsColumnDirective                           []string `json:"lighthouse-core/audits/csp-xss.js | columnDirective"`
			LighthouseCoreAuditsCspXssJsColumnSeverity                            []string `json:"lighthouse-core/audits/csp-xss.js | columnSeverity"`
			LighthouseCoreLibCspEvaluatorJsStrictDynamic                          []string `json:"lighthouse-core/lib/csp-evaluator.js | strictDynamic"`
			LighthouseCoreLibI18NI18NJsItemSeverityHigh                           []string `json:"lighthouse-core/lib/i18n/i18n.js | itemSeverityHigh"`
			LighthouseCoreAuditsManualPwaCrossBrowserJsTitle                      []string `json:"lighthouse-core/audits/manual/pwa-cross-browser.js | title"`
			LighthouseCoreAuditsManualPwaCrossBrowserJsDescription                []string `json:"lighthouse-core/audits/manual/pwa-cross-browser.js | description"`
			LighthouseCoreAuditsManualPwaPageTransitionsJsTitle                   []string `json:"lighthouse-core/audits/manual/pwa-page-transitions.js | title"`
			LighthouseCoreAuditsManualPwaPageTransitionsJsDescription             []string `json:"lighthouse-core/audits/manual/pwa-page-transitions.js | description"`
			LighthouseCoreAuditsManualPwaEachPageHasUrlJsTitle                    []string `json:"lighthouse-core/audits/manual/pwa-each-page-has-url.js | title"`
			LighthouseCoreAuditsManualPwaEachPageHasUrlJsDescription              []string `json:"lighthouse-core/audits/manual/pwa-each-page-has-url.js | description"`
			LighthouseCoreAuditsAccessibilityAccesskeysJsTitle                    []string `json:"lighthouse-core/audits/accessibility/accesskeys.js | title"`
			LighthouseCoreAuditsAccessibilityAccesskeysJsDescription              []string `json:"lighthouse-core/audits/accessibility/accesskeys.js | description"`
			LighthouseCoreAuditsAccessibilityAriaAllowedAttrJsTitle               []string `json:"lighthouse-core/audits/accessibility/aria-allowed-attr.js | title"`
			LighthouseCoreAuditsAccessibilityAriaAllowedAttrJsDescription         []string `json:"lighthouse-core/audits/accessibility/aria-allowed-attr.js | description"`
			LighthouseCoreAuditsAccessibilityAriaCommandNameJsTitle               []string `json:"lighthouse-core/audits/accessibility/aria-command-name.js | title"`
			LighthouseCoreAuditsAccessibilityAriaCommandNameJsDescription         []string `json:"lighthouse-core/audits/accessibility/aria-command-name.js | description"`
			LighthouseCoreAuditsAccessibilityAriaHiddenBodyJsTitle                []string `json:"lighthouse-core/audits/accessibility/aria-hidden-body.js | title"`
			LighthouseCoreAuditsAccessibilityAriaHiddenBodyJsDescription          []string `json:"lighthouse-core/audits/accessibility/aria-hidden-body.js | description"`
			LighthouseCoreAuditsAccessibilityAriaHiddenFocusJsTitle               []string `json:"lighthouse-core/audits/accessibility/aria-hidden-focus.js | title"`
			LighthouseCoreAuditsAccessibilityAriaHiddenFocusJsDescription         []string `json:"lighthouse-core/audits/accessibility/aria-hidden-focus.js | description"`
			LighthouseCoreAuditsAccessibilityAriaInputFieldNameJsTitle            []string `json:"lighthouse-core/audits/accessibility/aria-input-field-name.js | title"`
			LighthouseCoreAuditsAccessibilityAriaInputFieldNameJsDescription      []string `json:"lighthouse-core/audits/accessibility/aria-input-field-name.js | description"`
			LighthouseCoreAuditsAccessibilityAriaMeterNameJsTitle                 []string `json:"lighthouse-core/audits/accessibility/aria-meter-name.js | title"`
			LighthouseCoreAuditsAccessibilityAriaMeterNameJsDescription           []string `json:"lighthouse-core/audits/accessibility/aria-meter-name.js | description"`
			LighthouseCoreAuditsAccessibilityAriaProgressbarNameJsTitle           []string `json:"lighthouse-core/audits/accessibility/aria-progressbar-name.js | title"`
			LighthouseCoreAuditsAccessibilityAriaProgressbarNameJsDescription     []string `json:"lighthouse-core/audits/accessibility/aria-progressbar-name.js | description"`
			LighthouseCoreAuditsAccessibilityAriaRequiredAttrJsTitle              []string `json:"lighthouse-core/audits/accessibility/aria-required-attr.js | title"`
			LighthouseCoreAuditsAccessibilityAriaRequiredAttrJsDescription        []string `json:"lighthouse-core/audits/accessibility/aria-required-attr.js | description"`
			LighthouseCoreAuditsAccessibilityAriaRequiredChildrenJsFailureTitle   []string `json:"lighthouse-core/audits/accessibility/aria-required-children.js | failureTitle"`
			LighthouseCoreAuditsAccessibilityAriaRequiredChildrenJsDescription    []string `json:"lighthouse-core/audits/accessibility/aria-required-children.js | description"`
			LighthouseCoreLibI18NI18NJsColumnFailingElem                          []string `json:"lighthouse-core/lib/i18n/i18n.js | columnFailingElem"`
			LighthouseCoreAuditsAccessibilityAriaRequiredParentJsFailureTitle     []string `json:"lighthouse-core/audits/accessibility/aria-required-parent.js | failureTitle"`
			LighthouseCoreAuditsAccessibilityAriaRequiredParentJsDescription      []string `json:"lighthouse-core/audits/accessibility/aria-required-parent.js | description"`
			LighthouseCoreAuditsAccessibilityAriaRolesJsTitle                     []string `json:"lighthouse-core/audits/accessibility/aria-roles.js | title"`
			LighthouseCoreAuditsAccessibilityAriaRolesJsDescription               []string `json:"lighthouse-core/audits/accessibility/aria-roles.js | description"`
			LighthouseCoreAuditsAccessibilityAriaToggleFieldNameJsTitle           []string `json:"lighthouse-core/audits/accessibility/aria-toggle-field-name.js | title"`
			LighthouseCoreAuditsAccessibilityAriaToggleFieldNameJsDescription     []string `json:"lighthouse-core/audits/accessibility/aria-toggle-field-name.js | description"`
			LighthouseCoreAuditsAccessibilityAriaTooltipNameJsTitle               []string `json:"lighthouse-core/audits/accessibility/aria-tooltip-name.js | title"`
			LighthouseCoreAuditsAccessibilityAriaTooltipNameJsDescription         []string `json:"lighthouse-core/audits/accessibility/aria-tooltip-name.js | description"`
			LighthouseCoreAuditsAccessibilityAriaTreeitemNameJsTitle              []string `json:"lighthouse-core/audits/accessibility/aria-treeitem-name.js | title"`
			LighthouseCoreAuditsAccessibilityAriaTreeitemNameJsDescription        []string `json:"lighthouse-core/audits/accessibility/aria-treeitem-name.js | description"`
			LighthouseCoreAuditsAccessibilityAriaValidAttrValueJsTitle            []string `json:"lighthouse-core/audits/accessibility/aria-valid-attr-value.js | title"`
			LighthouseCoreAuditsAccessibilityAriaValidAttrValueJsDescription      []string `json:"lighthouse-core/audits/accessibility/aria-valid-attr-value.js | description"`
			LighthouseCoreAuditsAccessibilityAriaValidAttrJsTitle                 []string `json:"lighthouse-core/audits/accessibility/aria-valid-attr.js | title"`
			LighthouseCoreAuditsAccessibilityAriaValidAttrJsDescription           []string `json:"lighthouse-core/audits/accessibility/aria-valid-attr.js | description"`
			LighthouseCoreAuditsAccessibilityButtonNameJsTitle                    []string `json:"lighthouse-core/audits/accessibility/button-name.js | title"`
			LighthouseCoreAuditsAccessibilityButtonNameJsDescription              []string `json:"lighthouse-core/audits/accessibility/button-name.js | description"`
			LighthouseCoreAuditsAccessibilityBypassJsTitle                        []string `json:"lighthouse-core/audits/accessibility/bypass.js | title"`
			LighthouseCoreAuditsAccessibilityBypassJsDescription                  []string `json:"lighthouse-core/audits/accessibility/bypass.js | description"`
			LighthouseCoreAuditsAccessibilityColorContrastJsFailureTitle          []string `json:"lighthouse-core/audits/accessibility/color-contrast.js | failureTitle"`
			LighthouseCoreAuditsAccessibilityColorContrastJsDescription           []string `json:"lighthouse-core/audits/accessibility/color-contrast.js | description"`
			LighthouseCoreAuditsAccessibilityDefinitionListJsTitle                []string `json:"lighthouse-core/audits/accessibility/definition-list.js | title"`
			LighthouseCoreAuditsAccessibilityDefinitionListJsDescription          []string `json:"lighthouse-core/audits/accessibility/definition-list.js | description"`
			LighthouseCoreAuditsAccessibilityDlitemJsTitle                        []string `json:"lighthouse-core/audits/accessibility/dlitem.js | title"`
			LighthouseCoreAuditsAccessibilityDlitemJsDescription                  []string `json:"lighthouse-core/audits/accessibility/dlitem.js | description"`
			LighthouseCoreAuditsAccessibilityDocumentTitleJsTitle                 []string `json:"lighthouse-core/audits/accessibility/document-title.js | title"`
			LighthouseCoreAuditsAccessibilityDocumentTitleJsDescription           []string `json:"lighthouse-core/audits/accessibility/document-title.js | description"`
			LighthouseCoreAuditsAccessibilityDuplicateIdActiveJsTitle             []string `json:"lighthouse-core/audits/accessibility/duplicate-id-active.js | title"`
			LighthouseCoreAuditsAccessibilityDuplicateIdActiveJsDescription       []string `json:"lighthouse-core/audits/accessibility/duplicate-id-active.js | description"`
			LighthouseCoreAuditsAccessibilityDuplicateIdAriaJsTitle               []string `json:"lighthouse-core/audits/accessibility/duplicate-id-aria.js | title"`
			LighthouseCoreAuditsAccessibilityDuplicateIdAriaJsDescription         []string `json:"lighthouse-core/audits/accessibility/duplicate-id-aria.js | description"`
			LighthouseCoreAuditsAccessibilityFormFieldMultipleLabelsJsTitle       []string `json:"lighthouse-core/audits/accessibility/form-field-multiple-labels.js | title"`
			LighthouseCoreAuditsAccessibilityFormFieldMultipleLabelsJsDescription []string `json:"lighthouse-core/audits/accessibility/form-field-multiple-labels.js | description"`
			LighthouseCoreAuditsAccessibilityFrameTitleJsTitle                    []string `json:"lighthouse-core/audits/accessibility/frame-title.js | title"`
			LighthouseCoreAuditsAccessibilityFrameTitleJsDescription              []string `json:"lighthouse-core/audits/accessibility/frame-title.js | description"`
			LighthouseCoreAuditsAccessibilityHeadingOrderJsFailureTitle           []string `json:"lighthouse-core/audits/accessibility/heading-order.js | failureTitle"`
			LighthouseCoreAuditsAccessibilityHeadingOrderJsDescription            []string `json:"lighthouse-core/audits/accessibility/heading-order.js | description"`
			LighthouseCoreAuditsAccessibilityHtmlHasLangJsTitle                   []string `json:"lighthouse-core/audits/accessibility/html-has-lang.js | title"`
			LighthouseCoreAuditsAccessibilityHtmlHasLangJsDescription             []string `json:"lighthouse-core/audits/accessibility/html-has-lang.js | description"`
			LighthouseCoreAuditsAccessibilityHtmlLangValidJsTitle                 []string `json:"lighthouse-core/audits/accessibility/html-lang-valid.js | title"`
			LighthouseCoreAuditsAccessibilityHtmlLangValidJsDescription           []string `json:"lighthouse-core/audits/accessibility/html-lang-valid.js | description"`
			LighthouseCoreAuditsAccessibilityImageAltJsTitle                      []string `json:"lighthouse-core/audits/accessibility/image-alt.js | title"`
			LighthouseCoreAuditsAccessibilityImageAltJsDescription                []string `json:"lighthouse-core/audits/accessibility/image-alt.js | description"`
			LighthouseCoreAuditsAccessibilityInputImageAltJsTitle                 []string `json:"lighthouse-core/audits/accessibility/input-image-alt.js | title"`
			LighthouseCoreAuditsAccessibilityInputImageAltJsDescription           []string `json:"lighthouse-core/audits/accessibility/input-image-alt.js | description"`
			LighthouseCoreAuditsAccessibilityLabelJsTitle                         []string `json:"lighthouse-core/audits/accessibility/label.js | title"`
			LighthouseCoreAuditsAccessibilityLabelJsDescription                   []string `json:"lighthouse-core/audits/accessibility/label.js | description"`
			LighthouseCoreAuditsAccessibilityLinkNameJsTitle                      []string `json:"lighthouse-core/audits/accessibility/link-name.js | title"`
			LighthouseCoreAuditsAccessibilityLinkNameJsDescription                []string `json:"lighthouse-core/audits/accessibility/link-name.js | description"`
			LighthouseCoreAuditsAccessibilityListJsTitle                          []string `json:"lighthouse-core/audits/accessibility/list.js | title"`
			LighthouseCoreAuditsAccessibilityListJsDescription                    []string `json:"lighthouse-core/audits/accessibility/list.js | description"`
			LighthouseCoreAuditsAccessibilityListitemJsTitle                      []string `json:"lighthouse-core/audits/accessibility/listitem.js | title"`
			LighthouseCoreAuditsAccessibilityListitemJsDescription                []string `json:"lighthouse-core/audits/accessibility/listitem.js | description"`
			LighthouseCoreAuditsAccessibilityMetaRefreshJsTitle                   []string `json:"lighthouse-core/audits/accessibility/meta-refresh.js | title"`
			LighthouseCoreAuditsAccessibilityMetaRefreshJsDescription             []string `json:"lighthouse-core/audits/accessibility/meta-refresh.js | description"`
			LighthouseCoreAuditsAccessibilityMetaViewportJsTitle                  []string `json:"lighthouse-core/audits/accessibility/meta-viewport.js | title"`
			LighthouseCoreAuditsAccessibilityMetaViewportJsDescription            []string `json:"lighthouse-core/audits/accessibility/meta-viewport.js | description"`
			LighthouseCoreAuditsAccessibilityObjectAltJsTitle                     []string `json:"lighthouse-core/audits/accessibility/object-alt.js | title"`
			LighthouseCoreAuditsAccessibilityObjectAltJsDescription               []string `json:"lighthouse-core/audits/accessibility/object-alt.js | description"`
			LighthouseCoreAuditsAccessibilityTabindexJsTitle                      []string `json:"lighthouse-core/audits/accessibility/tabindex.js | title"`
			LighthouseCoreAuditsAccessibilityTabindexJsDescription                []string `json:"lighthouse-core/audits/accessibility/tabindex.js | description"`
			LighthouseCoreAuditsAccessibilityTdHeadersAttrJsTitle                 []string `json:"lighthouse-core/audits/accessibility/td-headers-attr.js | title"`
			LighthouseCoreAuditsAccessibilityTdHeadersAttrJsDescription           []string `json:"lighthouse-core/audits/accessibility/td-headers-attr.js | description"`
			LighthouseCoreAuditsAccessibilityThHasDataCellsJsTitle                []string `json:"lighthouse-core/audits/accessibility/th-has-data-cells.js | title"`
			LighthouseCoreAuditsAccessibilityThHasDataCellsJsDescription          []string `json:"lighthouse-core/audits/accessibility/th-has-data-cells.js | description"`
			LighthouseCoreAuditsAccessibilityValidLangJsTitle                     []string `json:"lighthouse-core/audits/accessibility/valid-lang.js | title"`
			LighthouseCoreAuditsAccessibilityValidLangJsDescription               []string `json:"lighthouse-core/audits/accessibility/valid-lang.js | description"`
			LighthouseCoreAuditsAccessibilityVideoCaptionJsTitle                  []string `json:"lighthouse-core/audits/accessibility/video-caption.js | title"`
			LighthouseCoreAuditsAccessibilityVideoCaptionJsDescription            []string `json:"lighthouse-core/audits/accessibility/video-caption.js | description"`
			LighthouseCoreAuditsByteEfficiencyUsesLongCacheTtlJsFailureTitle      []string `json:"lighthouse-core/audits/byte-efficiency/uses-long-cache-ttl.js | failureTitle"`
			LighthouseCoreAuditsByteEfficiencyUsesLongCacheTtlJsDescription       []string `json:"lighthouse-core/audits/byte-efficiency/uses-long-cache-ttl.js | description"`
			LighthouseCoreAuditsByteEfficiencyUsesLongCacheTtlJsDisplayValue      []struct {
				Values struct {
					ItemCount float64 `json:"itemCount"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/audits/byte-efficiency/uses-long-cache-ttl.js | displayValue"`
			LighthouseCoreLibI18NI18NJsColumnCacheTTL                       []string `json:"lighthouse-core/lib/i18n/i18n.js | columnCacheTTL"`
			LighthouseCoreAuditsByteEfficiencyTotalByteWeightJsTitle        []string `json:"lighthouse-core/audits/byte-efficiency/total-byte-weight.js | title"`
			LighthouseCoreAuditsByteEfficiencyTotalByteWeightJsDescription  []string `json:"lighthouse-core/audits/byte-efficiency/total-byte-weight.js | description"`
			LighthouseCoreAuditsByteEfficiencyTotalByteWeightJsDisplayValue []struct {
				Values struct {
					TotalBytes float64 `json:"totalBytes"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/audits/byte-efficiency/total-byte-weight.js | displayValue"`
			LighthouseCoreAuditsByteEfficiencyOffscreenImagesJsTitle               []string `json:"lighthouse-core/audits/byte-efficiency/offscreen-images.js | title"`
			LighthouseCoreAuditsByteEfficiencyOffscreenImagesJsDescription         []string `json:"lighthouse-core/audits/byte-efficiency/offscreen-images.js | description"`
			LighthouseCoreAuditsByteEfficiencyRenderBlockingResourcesJsTitle       []string `json:"lighthouse-core/audits/byte-efficiency/render-blocking-resources.js | title"`
			LighthouseCoreAuditsByteEfficiencyRenderBlockingResourcesJsDescription []string `json:"lighthouse-core/audits/byte-efficiency/render-blocking-resources.js | description"`
			LighthouseCoreLibI18NI18NJsDisplayValueMsSavings                       []struct {
				Values struct {
					WastedMs float64 `json:"wastedMs"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/lib/i18n/i18n.js | displayValueMsSavings"`
			LighthouseCoreAuditsByteEfficiencyUnminifiedCssJsTitle              []string `json:"lighthouse-core/audits/byte-efficiency/unminified-css.js | title"`
			LighthouseCoreAuditsByteEfficiencyUnminifiedCssJsDescription        []string `json:"lighthouse-core/audits/byte-efficiency/unminified-css.js | description"`
			LighthouseCoreAuditsByteEfficiencyUnminifiedJavascriptJsTitle       []string `json:"lighthouse-core/audits/byte-efficiency/unminified-javascript.js | title"`
			LighthouseCoreAuditsByteEfficiencyUnminifiedJavascriptJsDescription []string `json:"lighthouse-core/audits/byte-efficiency/unminified-javascript.js | description"`
			LighthouseCoreAuditsByteEfficiencyUnusedCssRulesJsTitle             []string `json:"lighthouse-core/audits/byte-efficiency/unused-css-rules.js | title"`
			LighthouseCoreAuditsByteEfficiencyUnusedCssRulesJsDescription       []string `json:"lighthouse-core/audits/byte-efficiency/unused-css-rules.js | description"`
			LighthouseCoreLibI18NI18NJsDisplayValueByteSavings                  []struct {
				Values struct {
					WastedBytes float64 `json:"wastedBytes"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/lib/i18n/i18n.js | displayValueByteSavings"`
			LighthouseCoreAuditsByteEfficiencyUnusedJavascriptJsTitle               []string `json:"lighthouse-core/audits/byte-efficiency/unused-javascript.js | title"`
			LighthouseCoreAuditsByteEfficiencyUnusedJavascriptJsDescription         []string `json:"lighthouse-core/audits/byte-efficiency/unused-javascript.js | description"`
			LighthouseCoreAuditsByteEfficiencyModernImageFormatsJsTitle             []string `json:"lighthouse-core/audits/byte-efficiency/modern-image-formats.js | title"`
			LighthouseCoreAuditsByteEfficiencyModernImageFormatsJsDescription       []string `json:"lighthouse-core/audits/byte-efficiency/modern-image-formats.js | description"`
			LighthouseCoreLibI18NI18NJsColumnResourceSize                           []string `json:"lighthouse-core/lib/i18n/i18n.js | columnResourceSize"`
			LighthouseCoreAuditsByteEfficiencyUsesOptimizedImagesJsTitle            []string `json:"lighthouse-core/audits/byte-efficiency/uses-optimized-images.js | title"`
			LighthouseCoreAuditsByteEfficiencyUsesOptimizedImagesJsDescription      []string `json:"lighthouse-core/audits/byte-efficiency/uses-optimized-images.js | description"`
			LighthouseCoreAuditsByteEfficiencyUsesTextCompressionJsTitle            []string `json:"lighthouse-core/audits/byte-efficiency/uses-text-compression.js | title"`
			LighthouseCoreAuditsByteEfficiencyUsesTextCompressionJsDescription      []string `json:"lighthouse-core/audits/byte-efficiency/uses-text-compression.js | description"`
			LighthouseCoreAuditsByteEfficiencyUsesResponsiveImagesJsTitle           []string `json:"lighthouse-core/audits/byte-efficiency/uses-responsive-images.js | title"`
			LighthouseCoreAuditsByteEfficiencyUsesResponsiveImagesJsDescription     []string `json:"lighthouse-core/audits/byte-efficiency/uses-responsive-images.js | description"`
			LighthouseCoreAuditsByteEfficiencyEfficientAnimatedContentJsTitle       []string `json:"lighthouse-core/audits/byte-efficiency/efficient-animated-content.js | title"`
			LighthouseCoreAuditsByteEfficiencyEfficientAnimatedContentJsDescription []string `json:"lighthouse-core/audits/byte-efficiency/efficient-animated-content.js | description"`
			LighthouseCoreAuditsByteEfficiencyDuplicatedJavascriptJsTitle           []string `json:"lighthouse-core/audits/byte-efficiency/duplicated-javascript.js | title"`
			LighthouseCoreAuditsByteEfficiencyDuplicatedJavascriptJsDescription     []string `json:"lighthouse-core/audits/byte-efficiency/duplicated-javascript.js | description"`
			LighthouseCoreAuditsByteEfficiencyLegacyJavascriptJsTitle               []string `json:"lighthouse-core/audits/byte-efficiency/legacy-javascript.js | title"`
			LighthouseCoreAuditsByteEfficiencyLegacyJavascriptJsDescription         []string `json:"lighthouse-core/audits/byte-efficiency/legacy-javascript.js | description"`
			LighthouseCoreAuditsDobetterwebDoctypeJsTitle                           []string `json:"lighthouse-core/audits/dobetterweb/doctype.js | title"`
			LighthouseCoreAuditsDobetterwebDoctypeJsDescription                     []string `json:"lighthouse-core/audits/dobetterweb/doctype.js | description"`
			LighthouseCoreAuditsDobetterwebCharsetJsTitle                           []string `json:"lighthouse-core/audits/dobetterweb/charset.js | title"`
			LighthouseCoreAuditsDobetterwebCharsetJsDescription                     []string `json:"lighthouse-core/audits/dobetterweb/charset.js | description"`
			LighthouseCoreAuditsDobetterwebDomSizeJsFailureTitle                    []string `json:"lighthouse-core/audits/dobetterweb/dom-size.js | failureTitle"`
			LighthouseCoreAuditsDobetterwebDomSizeJsDescription                     []string `json:"lighthouse-core/audits/dobetterweb/dom-size.js | description"`
			LighthouseCoreAuditsDobetterwebDomSizeJsDisplayValue                    []struct {
				Values struct {
					ItemCount float64 `json:"itemCount"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/audits/dobetterweb/dom-size.js | displayValue"`
			LighthouseCoreAuditsDobetterwebDomSizeJsColumnStatistic                   []string `json:"lighthouse-core/audits/dobetterweb/dom-size.js | columnStatistic"`
			LighthouseCoreAuditsDobetterwebDomSizeJsColumnValue                       []string `json:"lighthouse-core/audits/dobetterweb/dom-size.js | columnValue"`
			LighthouseCoreAuditsDobetterwebDomSizeJsStatisticDOMElements              []string `json:"lighthouse-core/audits/dobetterweb/dom-size.js | statisticDOMElements"`
			LighthouseCoreAuditsDobetterwebDomSizeJsStatisticDOMDepth                 []string `json:"lighthouse-core/audits/dobetterweb/dom-size.js | statisticDOMDepth"`
			LighthouseCoreAuditsDobetterwebDomSizeJsStatisticDOMWidth                 []string `json:"lighthouse-core/audits/dobetterweb/dom-size.js | statisticDOMWidth"`
			LighthouseCoreAuditsDobetterwebGeolocationOnStartJsTitle                  []string `json:"lighthouse-core/audits/dobetterweb/geolocation-on-start.js | title"`
			LighthouseCoreAuditsDobetterwebGeolocationOnStartJsDescription            []string `json:"lighthouse-core/audits/dobetterweb/geolocation-on-start.js | description"`
			LighthouseCoreAuditsDobetterwebInspectorIssuesJsTitle                     []string `json:"lighthouse-core/audits/dobetterweb/inspector-issues.js | title"`
			LighthouseCoreAuditsDobetterwebInspectorIssuesJsDescription               []string `json:"lighthouse-core/audits/dobetterweb/inspector-issues.js | description"`
			LighthouseCoreAuditsDobetterwebNoDocumentWriteJsTitle                     []string `json:"lighthouse-core/audits/dobetterweb/no-document-write.js | title"`
			LighthouseCoreAuditsDobetterwebNoDocumentWriteJsDescription               []string `json:"lighthouse-core/audits/dobetterweb/no-document-write.js | description"`
			LighthouseCoreAuditsDobetterwebNoVulnerableLibrariesJsTitle               []string `json:"lighthouse-core/audits/dobetterweb/no-vulnerable-libraries.js | title"`
			LighthouseCoreAuditsDobetterwebNoVulnerableLibrariesJsDescription         []string `json:"lighthouse-core/audits/dobetterweb/no-vulnerable-libraries.js | description"`
			LighthouseCoreAuditsDobetterwebJsLibrariesJsTitle                         []string `json:"lighthouse-core/audits/dobetterweb/js-libraries.js | title"`
			LighthouseCoreAuditsDobetterwebJsLibrariesJsDescription                   []string `json:"lighthouse-core/audits/dobetterweb/js-libraries.js | description"`
			LighthouseCoreLibI18NI18NJsColumnName                                     []string `json:"lighthouse-core/lib/i18n/i18n.js | columnName"`
			LighthouseCoreAuditsDobetterwebJsLibrariesJsColumnVersion                 []string `json:"lighthouse-core/audits/dobetterweb/js-libraries.js | columnVersion"`
			LighthouseCoreAuditsDobetterwebNotificationOnStartJsTitle                 []string `json:"lighthouse-core/audits/dobetterweb/notification-on-start.js | title"`
			LighthouseCoreAuditsDobetterwebNotificationOnStartJsDescription           []string `json:"lighthouse-core/audits/dobetterweb/notification-on-start.js | description"`
			LighthouseCoreAuditsDobetterwebPasswordInputsCanBePastedIntoJsTitle       []string `json:"lighthouse-core/audits/dobetterweb/password-inputs-can-be-pasted-into.js | title"`
			LighthouseCoreAuditsDobetterwebPasswordInputsCanBePastedIntoJsDescription []string `json:"lighthouse-core/audits/dobetterweb/password-inputs-can-be-pasted-into.js | description"`
			LighthouseCoreAuditsDobetterwebUsesHttp2JsTitle                           []string `json:"lighthouse-core/audits/dobetterweb/uses-http2.js | title"`
			LighthouseCoreAuditsDobetterwebUsesHttp2JsDescription                     []string `json:"lighthouse-core/audits/dobetterweb/uses-http2.js | description"`
			LighthouseCoreAuditsDobetterwebUsesPassiveEventListenersJsFailureTitle    []string `json:"lighthouse-core/audits/dobetterweb/uses-passive-event-listeners.js | failureTitle"`
			LighthouseCoreAuditsDobetterwebUsesPassiveEventListenersJsDescription     []string `json:"lighthouse-core/audits/dobetterweb/uses-passive-event-listeners.js | description"`
			LighthouseCoreAuditsSeoMetaDescriptionJsTitle                             []string `json:"lighthouse-core/audits/seo/meta-description.js | title"`
			LighthouseCoreAuditsSeoMetaDescriptionJsDescription                       []string `json:"lighthouse-core/audits/seo/meta-description.js | description"`
			LighthouseCoreAuditsSeoHttpStatusCodeJsTitle                              []string `json:"lighthouse-core/audits/seo/http-status-code.js | title"`
			LighthouseCoreAuditsSeoHttpStatusCodeJsDescription                        []string `json:"lighthouse-core/audits/seo/http-status-code.js | description"`
			LighthouseCoreAuditsSeoFontSizeJsTitle                                    []string `json:"lighthouse-core/audits/seo/font-size.js | title"`
			LighthouseCoreAuditsSeoFontSizeJsDescription                              []string `json:"lighthouse-core/audits/seo/font-size.js | description"`
			LighthouseCoreAuditsSeoFontSizeJsDisplayValue                             []struct {
				Values struct {
					DecimalProportion float64 `json:"decimalProportion"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/audits/seo/font-size.js | displayValue"`
			LighthouseCoreAuditsSeoFontSizeJsColumnSelector            []string `json:"lighthouse-core/audits/seo/font-size.js | columnSelector"`
			LighthouseCoreAuditsSeoFontSizeJsColumnPercentPageText     []string `json:"lighthouse-core/audits/seo/font-size.js | columnPercentPageText"`
			LighthouseCoreAuditsSeoFontSizeJsColumnFontSize            []string `json:"lighthouse-core/audits/seo/font-size.js | columnFontSize"`
			LighthouseCoreAuditsSeoFontSizeJsLegibleText               []string `json:"lighthouse-core/audits/seo/font-size.js | legibleText"`
			LighthouseCoreAuditsSeoLinkTextJsTitle                     []string `json:"lighthouse-core/audits/seo/link-text.js | title"`
			LighthouseCoreAuditsSeoLinkTextJsDescription               []string `json:"lighthouse-core/audits/seo/link-text.js | description"`
			LighthouseCoreAuditsSeoCrawlableAnchorsJsFailureTitle      []string `json:"lighthouse-core/audits/seo/crawlable-anchors.js | failureTitle"`
			LighthouseCoreAuditsSeoCrawlableAnchorsJsDescription       []string `json:"lighthouse-core/audits/seo/crawlable-anchors.js | description"`
			LighthouseCoreAuditsSeoCrawlableAnchorsJsColumnFailingLink []string `json:"lighthouse-core/audits/seo/crawlable-anchors.js | columnFailingLink"`
			LighthouseCoreAuditsSeoIsCrawlableJsTitle                  []string `json:"lighthouse-core/audits/seo/is-crawlable.js | title"`
			LighthouseCoreAuditsSeoIsCrawlableJsDescription            []string `json:"lighthouse-core/audits/seo/is-crawlable.js | description"`
			LighthouseCoreAuditsSeoRobotsTxtJsTitle                    []string `json:"lighthouse-core/audits/seo/robots-txt.js | title"`
			LighthouseCoreAuditsSeoRobotsTxtJsDescription              []string `json:"lighthouse-core/audits/seo/robots-txt.js | description"`
			LighthouseCoreAuditsSeoTapTargetsJsTitle                   []string `json:"lighthouse-core/audits/seo/tap-targets.js | title"`
			LighthouseCoreAuditsSeoTapTargetsJsDescription             []string `json:"lighthouse-core/audits/seo/tap-targets.js | description"`
			LighthouseCoreAuditsSeoTapTargetsJsDisplayValue            []struct {
				Values struct {
					DecimalProportion float64 `json:"decimalProportion"`
				} `json:"values"`
				Path string `json:"path"`
			} `json:"lighthouse-core/audits/seo/tap-targets.js | displayValue"`
			LighthouseCoreAuditsSeoHreflangJsTitle                                  []string `json:"lighthouse-core/audits/seo/hreflang.js | title"`
			LighthouseCoreAuditsSeoHreflangJsDescription                            []string `json:"lighthouse-core/audits/seo/hreflang.js | description"`
			LighthouseCoreAuditsSeoPluginsJsTitle                                   []string `json:"lighthouse-core/audits/seo/plugins.js | title"`
			LighthouseCoreAuditsSeoPluginsJsDescription                             []string `json:"lighthouse-core/audits/seo/plugins.js | description"`
			LighthouseCoreAuditsSeoCanonicalJsTitle                                 []string `json:"lighthouse-core/audits/seo/canonical.js | title"`
			LighthouseCoreAuditsSeoCanonicalJsDescription                           []string `json:"lighthouse-core/audits/seo/canonical.js | description"`
			LighthouseCoreAuditsSeoManualStructuredDataJsTitle                      []string `json:"lighthouse-core/audits/seo/manual/structured-data.js | title"`
			LighthouseCoreAuditsSeoManualStructuredDataJsDescription                []string `json:"lighthouse-core/audits/seo/manual/structured-data.js | description"`
			LighthouseCoreConfigDefaultConfigJsPerformanceCategoryTitle             []string `json:"lighthouse-core/config/default-config.js | performanceCategoryTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YCategoryTitle                    []string `json:"lighthouse-core/config/default-config.js | a11yCategoryTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YCategoryDescription              []string `json:"lighthouse-core/config/default-config.js | a11yCategoryDescription"`
			LighthouseCoreConfigDefaultConfigJsA11YCategoryManualDescription        []string `json:"lighthouse-core/config/default-config.js | a11yCategoryManualDescription"`
			LighthouseCoreConfigDefaultConfigJsBestPracticesCategoryTitle           []string `json:"lighthouse-core/config/default-config.js | bestPracticesCategoryTitle"`
			LighthouseCoreConfigDefaultConfigJsSeoCategoryTitle                     []string `json:"lighthouse-core/config/default-config.js | seoCategoryTitle"`
			LighthouseCoreConfigDefaultConfigJsSeoCategoryDescription               []string `json:"lighthouse-core/config/default-config.js | seoCategoryDescription"`
			LighthouseCoreConfigDefaultConfigJsSeoCategoryManualDescription         []string `json:"lighthouse-core/config/default-config.js | seoCategoryManualDescription"`
			LighthouseCoreConfigDefaultConfigJsPwaCategoryTitle                     []string `json:"lighthouse-core/config/default-config.js | pwaCategoryTitle"`
			LighthouseCoreConfigDefaultConfigJsPwaCategoryDescription               []string `json:"lighthouse-core/config/default-config.js | pwaCategoryDescription"`
			LighthouseCoreConfigDefaultConfigJsPwaCategoryManualDescription         []string `json:"lighthouse-core/config/default-config.js | pwaCategoryManualDescription"`
			LighthouseCoreConfigDefaultConfigJsMetricGroupTitle                     []string `json:"lighthouse-core/config/default-config.js | metricGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsLoadOpportunitiesGroupTitle          []string `json:"lighthouse-core/config/default-config.js | loadOpportunitiesGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsLoadOpportunitiesGroupDescription    []string `json:"lighthouse-core/config/default-config.js | loadOpportunitiesGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsBudgetsGroupTitle                    []string `json:"lighthouse-core/config/default-config.js | budgetsGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsBudgetsGroupDescription              []string `json:"lighthouse-core/config/default-config.js | budgetsGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsDiagnosticsGroupTitle                []string `json:"lighthouse-core/config/default-config.js | diagnosticsGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsDiagnosticsGroupDescription          []string `json:"lighthouse-core/config/default-config.js | diagnosticsGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsPwaInstallableGroupTitle             []string `json:"lighthouse-core/config/default-config.js | pwaInstallableGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsPwaOptimizedGroupTitle               []string `json:"lighthouse-core/config/default-config.js | pwaOptimizedGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YBestPracticesGroupTitle          []string `json:"lighthouse-core/config/default-config.js | a11yBestPracticesGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YBestPracticesGroupDescription    []string `json:"lighthouse-core/config/default-config.js | a11yBestPracticesGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsA11YColorContrastGroupTitle          []string `json:"lighthouse-core/config/default-config.js | a11yColorContrastGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YColorContrastGroupDescription    []string `json:"lighthouse-core/config/default-config.js | a11yColorContrastGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsA11YNamesLabelsGroupTitle            []string `json:"lighthouse-core/config/default-config.js | a11yNamesLabelsGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YNamesLabelsGroupDescription      []string `json:"lighthouse-core/config/default-config.js | a11yNamesLabelsGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsA11YNavigationGroupTitle             []string `json:"lighthouse-core/config/default-config.js | a11yNavigationGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YNavigationGroupDescription       []string `json:"lighthouse-core/config/default-config.js | a11yNavigationGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsA11YAriaGroupTitle                   []string `json:"lighthouse-core/config/default-config.js | a11yAriaGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YAriaGroupDescription             []string `json:"lighthouse-core/config/default-config.js | a11yAriaGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsA11YLanguageGroupTitle               []string `json:"lighthouse-core/config/default-config.js | a11yLanguageGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YLanguageGroupDescription         []string `json:"lighthouse-core/config/default-config.js | a11yLanguageGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsA11YAudioVideoGroupTitle             []string `json:"lighthouse-core/config/default-config.js | a11yAudioVideoGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YAudioVideoGroupDescription       []string `json:"lighthouse-core/config/default-config.js | a11yAudioVideoGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsA11YTablesListsVideoGroupTitle       []string `json:"lighthouse-core/config/default-config.js | a11yTablesListsVideoGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsA11YTablesListsVideoGroupDescription []string `json:"lighthouse-core/config/default-config.js | a11yTablesListsVideoGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsSeoMobileGroupTitle                  []string `json:"lighthouse-core/config/default-config.js | seoMobileGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsSeoMobileGroupDescription            []string `json:"lighthouse-core/config/default-config.js | seoMobileGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsSeoContentGroupTitle                 []string `json:"lighthouse-core/config/default-config.js | seoContentGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsSeoContentGroupDescription           []string `json:"lighthouse-core/config/default-config.js | seoContentGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsSeoCrawlingGroupTitle                []string `json:"lighthouse-core/config/default-config.js | seoCrawlingGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsSeoCrawlingGroupDescription          []string `json:"lighthouse-core/config/default-config.js | seoCrawlingGroupDescription"`
			LighthouseCoreConfigDefaultConfigJsBestPracticesTrustSafetyGroupTitle   []string `json:"lighthouse-core/config/default-config.js | bestPracticesTrustSafetyGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsBestPracticesUXGroupTitle            []string `json:"lighthouse-core/config/default-config.js | bestPracticesUXGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsBestPracticesBrowserCompatGroupTitle []string `json:"lighthouse-core/config/default-config.js | bestPracticesBrowserCompatGroupTitle"`
			LighthouseCoreConfigDefaultConfigJsBestPracticesGeneralGroupTitle       []string `json:"lighthouse-core/config/default-config.js | bestPracticesGeneralGroupTitle"`
		} `json:"icuMessagePaths"`
	} `json:"i18n"`
}
