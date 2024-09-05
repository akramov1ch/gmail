package handlers

import (
    "context"
    "contact-management-service/internal/contacts"
    "contact-management-service/proto"
)

type ContactManagementService struct {
    proto.UnimplementedContactManagementServiceServer
}

func (s *ContactManagementService) AddContact(ctx context.Context, req *proto.AddContactRequest) (*proto.AddContactResponse, error) {
    contactId, err := contacts.AddContact(req.UserId, req.Name, req.Email, req.PhoneNumber)
    if err != nil {
        return nil, err
    }
    return &proto.AddContactResponse{ContactId: contactId}, nil
}

func (s *ContactManagementService) UpdateContact(ctx context.Context, req *proto.UpdateContactRequest) (*proto.UpdateContactResponse, error) {
    err := contacts.UpdateContact(req.ContactId, req.Name, req.Email, req.PhoneNumber)
    if err != nil {
        return nil, err
    }
    return &proto.UpdateContactResponse{Status: "Success"}, nil
}

func (s *ContactManagementService) DeleteContact(ctx context.Context, req *proto.DeleteContactRequest) (*proto.DeleteContactResponse, error) {
    err := contacts.DeleteContact(req.ContactId)
    if err != nil {
        return nil, err
    }
    return &proto.DeleteContactResponse{Status: "Success"}, nil
}

func (s *ContactManagementService) GetContacts(ctx context.Context, req *proto.GetContactsRequest) (*proto.GetContactsResponse, error) {
    contacts, err := contacts.GetContacts(req.UserId)
    if err != nil {
        return nil, err
    }

    var protoContacts []*proto.Contact
    for _, c := range contacts {
        protoContacts = append(protoContacts, &proto.Contact{
            ContactId:   c.ContactId,
            Name:        c.Name,
            Email:       c.Email,
            PhoneNumber: c.PhoneNumber,
        })
    }

    return &proto.GetContactsResponse{Contacts: protoContacts}, nil
}

func (s *ContactManagementService) SearchContacts(ctx context.Context, req *proto.SearchContactsRequest) (*proto.SearchContactsResponse, error) {
    contacts, err := contacts.SearchContacts(req.UserId, req.Query)
    if err != nil {
        return nil, err
    }

    var protoContacts []*proto.Contact
    for _, c := range contacts {
        protoContacts = append(protoContacts, &proto.Contact{
            ContactId:   c.ContactId,
            Name:        c.Name,
            Email:       c.Email,
            PhoneNumber: c.PhoneNumber,
        })
    }

    return &proto.SearchContactsResponse{Contacts: protoContacts}, nil
}
