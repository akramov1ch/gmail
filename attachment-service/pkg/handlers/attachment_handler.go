package handlers

import (
	"context"
	"database/sql"
	"os"

	"attachment-service/internal/db"
	"attachment-service/internal/storage"
	"attachment-service/pkg/models"
	"attachment-service/proto"
)

type AttachmentService struct {
	proto.UnimplementedAttachmentServiceServer
}

func (s *AttachmentService) UploadAttachment(ctx context.Context, req *proto.UploadAttachmentRequest) (*proto.UploadAttachmentResponse, error) {
	filename := req.Filename
	contentType := req.ContentType

	fileURL, err := storage.UploadFile(os.File{}, filename, contentType)
	if err != nil {
		return nil, err
	}

	_, err = db.DB.Exec(
		"INSERT INTO attachments (filename, content_type, url) VALUES ($1, $2, $3)",
		filename, contentType, fileURL,
	)
	if err != nil {
		return nil, err
	}

	return &proto.UploadAttachmentResponse{Url: fileURL}, nil
}

func (s *AttachmentService) GetAttachment(ctx context.Context, req *proto.GetAttachmentRequest) (*proto.GetAttachmentResponse, error) {
	attachment := models.Attachment{}
	err := db.DB.QueryRow("SELECT filename, url FROM attachments WHERE id = $1", req.AttachmentId).Scan(&attachment.Filename, &attachment.URL)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &proto.GetAttachmentResponse{
		Filename: attachment.Filename,
		Url:      attachment.URL,
	}, nil
}
