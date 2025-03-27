resource "aws_db_subnet_group" "rds_subnet_group" {
  name = "${var.project_name}-rds-subnet"
  subnet_ids = [
    aws_subnet.private_a.id,
    aws_subnet.private_c.id
  ]

  tags = {
    Name    = "${var.project_name}-rds-subnet"
    Project = var.project_name
  }
}

resource "aws_security_group" "rds_sg" {
  name        = "${var.project_name}-rds-sg"
  description = "Security group for RDS instance"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port = 5432
    to_port   = 5432
    protocol  = "tcp"
    security_groups = [aws_security_group.api_service_sg.id]
  }

  egress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name    = "${var.project_name}-rds-sg"
    Project = var.project_name
  }
}

resource "aws_db_instance" "db" {
  identifier           = "${var.project_name}-postgres"
  engine               = "postgres"
  engine_version       = "17"
  instance_class       = "db.t3.micro"
  allocated_storage    = 20
  storage_type         = "gp2"
  username             = var.db_username
  password             = var.db_password
  db_name              = var.db_name
  publicly_accessible  = false
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  db_subnet_group_name = aws_db_subnet_group.rds_subnet_group.name

  skip_final_snapshot = true

  tags = {
    Name    = "${var.project_name}-postgres"
    Project = var.project_name
  }
}
