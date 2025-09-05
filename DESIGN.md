# Hexagonal Architecture Demonstration with Nix

## Overview

This project demonstrates **Hexagonal Architecture** (Ports and Adapters) principles using **Nix** for dependency management and application composition. The goal is to showcase how the same business logic can work with different user interfaces and data persistence layers through clean architectural boundaries.

## Architecture Concept

### Core Domain (The Hexagon)
- **Pure business logic** with no external dependencies
- **Domain entities** and **use cases** that represent the problem space
- **Technology-agnostic** - can work with any infrastructure

### Ports (Interfaces)
- **Primary Ports**: Interfaces that drive the application (e.g., TaskService)
- **Secondary Ports**: Interfaces the application depends on (e.g., TaskRepository)

### Adapters (The Outer Ring)
- **Primary Adapters**: User interfaces that call into the core (CLI, Web, TUI)
- **Secondary Adapters**: Infrastructure implementations (File storage, Database, etc.)

## Nix's Role

Nix enables us to:
1. **Compose different adapter combinations** as separate packages
2. **Inject dependencies** at build time through Nix expressions
3. **Create multiple application variants** from the same codebase
4. **Ensure reproducible builds** across different configurations

## Example Application: User Profile System

A simple user profile system that demonstrates hexagonal architecture with authentication and profile management:

### Core Domain
- **User Management**: Create, authenticate, and manage user accounts
- **Profile System**: Store and retrieve user profiles with text messages
- **Authentication**: Login/logout functionality with session management
- **Social Features**: View other users' profiles (basic social network)

### Primary Adapters (UI)
- **CLI**: Command-line interface for user registration, login, profile management
- **TUI**: Interactive terminal interface with menus and forms using ratatui
- **Web**: HTTP server with REST API and simple web frontend for browser access

### Secondary Adapters (Storage)
- **JSON File**: Simple file-based storage (`users.json`) for development/testing
- **PostgreSQL**: Full database with proper schemas, indexing, and ACID transactions

### Authentication & Session Management
- **File-based**: Sessions stored in JSON alongside user data
- **Database**: Sessions table with proper expiration and cleanup

### Example Application Variants

**Development Setup:**
- `user-tui-json`: Interactive TUI + JSON file storage
- `user-cli-json`: CLI commands + JSON file storage

**Production Setup:**
- `user-web-postgres`: Web interface + PostgreSQL database
- `user-tui-postgres`: TUI interface + PostgreSQL (admin tool)

### Nickel Configuration Management
Nickel manages the composition and environment-specific settings:
- Database connection strings and pool sizes
- File paths and formats
- Web server ports and CORS settings
- Authentication token expiration times
- UI-specific configurations (TUI themes, CLI help text)

This demonstrates how **the same user management business logic** can be deployed as a simple local tool (TUI+JSON) for personal use or as a scalable web service (Web+PostgreSQL) while maintaining clean architectural boundaries.
