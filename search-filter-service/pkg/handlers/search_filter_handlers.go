package handlers

import (
    "context"
    "search-service/internal/search"
    "search-service/internal/filter"
    "search-service/proto"
)

type SearchFilterService struct {
    proto.UnimplementedSearchFilterServiceServer
}

func (s *SearchFilterService) SearchEmails(ctx context.Context, req *proto.SearchEmailsRequest) (*proto.SearchEmailsResponse, error) {
    emails, err := search.SearchEmails(req.UserId, req.Query, req.Labels)
    if err != nil {
        return nil, err
    }

    var protoEmails []*proto.Email
    for _, email := range emails {
        protoEmails = append(protoEmails, &proto.Email{
            EmailId:  email.EmailId,
            Sender:   email.Sender,
            Subject:  email.Subject,
            Preview:  email.Preview,
        })
    }

    return &proto.SearchEmailsResponse{Emails: protoEmails}, nil
}

func (s *SearchFilterService) CreateFilter(ctx context.Context, req *proto.CreateFilterRequest) (*proto.CreateFilterResponse, error) {
    filterId, err := filter.CreateFilter(req.UserId, req.Sender, req.SubjectContains, req.MoveToFolder)
    if err != nil {
        return nil, err
    }
    return &proto.CreateFilterResponse{FilterId: filterId}, nil
}

func (s *SearchFilterService) ApplyFilters(ctx context.Context, req *proto.ApplyFiltersRequest) (*proto.ApplyFiltersResponse, error) {
    filteredEmails, err := filter.ApplyFilters(req.UserId)
    if err != nil {
        return nil, err
    }

    var protoFilteredEmails []*proto.FilteredEmail
    for _, fe := range filteredEmails {
        protoFilteredEmails = append(protoFilteredEmails, &proto.FilteredEmail{
            EmailId: fe.EmailId,
            Folder:  fe.Folder,
        })
    }

    return &proto.ApplyFiltersResponse{FilteredEmails: protoFilteredEmails}, nil
}
