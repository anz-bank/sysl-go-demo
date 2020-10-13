package handlers

import (
	"context"

	util "github.com/anz-bank/sysl-template/internal"
	"github.com/anz-bank/sysl-template/internal/gen/flickr"
)

// GetRestListRead reads photos from Flickr
func GetRestListRead(ctx context.Context,
	getRestListRequest *flickr.GetRestListRequest,
	client flickr.GetRestListClient) (*flickr.PhotoResource, error) {
	util.SetJSONResponseContentType(ctx)

	if *getRestListRequest.Method == "flickr.photos.search" {
		switch *getRestListRequest.Tags {
		case "dog":
			return &flickr.PhotoResource{
				Page:    util.NewFloat64(1),
				Pages:   util.NewFloat64(1),
				Perpage: util.NewFloat64(3),
				Photos: []flickr.Photo{
					{
						ID: util.NewString("white golden retriever"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: util.NewString("dog"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: util.NewString("http://www.example.com/dog/wgr"),
								},
							},
						},
					},
					{
						ID: util.NewString("I am JJ's dog Bingo"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: util.NewString("dog"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: util.NewString("http://www.example.com/dog/bingo"),
								},
							},
						},
					},
				},
				Total: util.NewFloat64(2),
			}, nil
		case "cat":
			return &flickr.PhotoResource{
				Page:    util.NewFloat64(1),
				Pages:   util.NewFloat64(1),
				Perpage: util.NewFloat64(3),
				Photos: []flickr.Photo{
					{
						ID: util.NewString("White cat"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: util.NewString("cat"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: util.NewString("http://www.example.com/cat/wc"),
								},
							},
						},
					},
					{
						ID: util.NewString("Black cat"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: util.NewString("cat"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: util.NewString("http://www.example.com/cat/bc"),
								},
							},
						},
					},
				},
				Total: util.NewFloat64(2),
			}, nil
		case "rabbit":
			return &flickr.PhotoResource{
				Page:    util.NewFloat64(1),
				Pages:   util.NewFloat64(1),
				Perpage: util.NewFloat64(3),
				Photos: []flickr.Photo{
					{
						ID: util.NewString("I am peter rabbit"),
						Tags: &flickr.Photo_tags{
							Tag: []flickr.Tag{
								{
									Raw: util.NewString("rabbit"),
								},
							},
						},
						Urls: &flickr.Photo_urls{
							URL: []flickr.URL{
								{
									Type: util.NewString("http://www.example.com/rabbit/pr"),
								},
							},
						},
					},
				},
				Total: util.NewFloat64(2),
			}, nil
		}
	}

	return nil, nil
}
